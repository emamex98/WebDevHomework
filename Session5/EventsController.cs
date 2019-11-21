// ESTE ES UN FRAGMENTO DEL CODIGO QUE HICE PARA LA API DEL CHANGEMAKER DAY

using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using AutoMapper;
using EventsApi.Domain.Entities;
using EventsApi.Domain.Models;
using EventsApi.Services;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Serilog;

namespace EventsApi.Controllers
{
    [AllowAnonymous]
    [Route("api/events")]
    [ApiController]

    public class EventsController : Controller
    {
        private readonly EventServices _eventServices;
        private readonly IMapper _mapper;

        public EventsController(EventServices eventServices, IMapper mapper)
        {
            _eventServices = eventServices;
            _mapper = mapper;
        }

        [HttpGet]
        public async Task<IActionResult> GetAllEvents()
        {
            var claims = Request.HttpContext.User;

            Log.Information("EventsController: HTTP GET events/");
            var result = await _eventServices.ListEventsAsync();
            var items = _mapper.Map<IEnumerable<Event>, IEnumerable<EventReadDto>>(result);

            foreach (var item in items)
            {
                String hrefUrl = $"{ Request.Scheme.ToString() }://{ Request.Host.ToString() }{ Request.Path.ToString() }/{ item.EventId.ToString() }/atendees";
                item.Links.Add(new LinkReadDto { Href = hrefUrl, Ref = "atendees", Type = "GET" });
            }

            return Ok(new { items });
        }

        [HttpGet("{eventId}")]
        public async Task<IActionResult> GetEventByID(int eventId)
        {
            Log.Information("EventsController: HTTP GET events/{eventId}", eventId);
            var result = await _eventServices.GetEventAsync(eventId);

            if (result == null)
                return NotFound();

            var item = _mapper.Map<Event, EventReadDto>(result);

            String hrefUrl = $"{ Request.Scheme.ToString() }://{ Request.Host.ToString() }{ Request.Path.ToString() }/atendees";
            item.Links.Add(new LinkReadDto { Href = hrefUrl, Ref = "atendees", Type = "GET" });

            return Ok(item);
        }

        [HttpPost]
        [ProducesResponseType(201)]
        public async Task<IActionResult> PostNewEvent([FromBody] EventWriteDto resource)
        {
            Log.Information("EventsController: HTTP POST events/");
            var newEvent = _mapper.Map<EventWriteDto, Event>(resource);
            var result = await _eventServices.AddEventAsync(newEvent);

            var createdObject = _mapper.Map<Event, EventReadDto>(result);
            String createdObjectURL = $"{ Request.Scheme.ToString() }://{ Request.Host.ToString() }{ Request.Path.ToString() }/{ createdObject.EventId.ToString() }";

            return Created(createdObjectURL, createdObject);
        }

        [HttpPut("{eventId}")]
        [ProducesResponseType(204)]
        public async Task<IActionResult> OverwriteEvent(int eventId, [FromBody] EventWriteDto resource)
        {
            Log.Information("EventsController: HTTP PUT events/{eventId}", eventId);
            var existingEvent = await _eventServices.GetEventAsync(eventId);

            if (existingEvent == null)
                return NotFound();

            var newEvent = _mapper.Map<EventWriteDto, Event>(resource);
            await _eventServices.OverwriteEventAsync(existingEvent, newEvent);

            return NoContent();
        }

        [HttpDelete("{eventId}")]
        [ProducesResponseType(204)]
        public async Task<IActionResult> DeleteEvent(int eventId)
        {
            Log.Information("EventssController: HTTP DELETE events/{eventId}", eventId);
            var existingEvent = await _eventServices.GetEventAsync(eventId);

            if (existingEvent == null)
                return NotFound();

            await _eventServices.DeleteEventAsync(existingEvent);

            return NoContent();
        }

        [AllowAnonymous]
        [HttpOptions]
        public IActionResult GetOptions()
        {
            return Ok();
        }
    }
}
