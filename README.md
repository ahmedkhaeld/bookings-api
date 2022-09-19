# Bookings-API
### Business 
The business is as we have a bed & breakfast hotel, that has room to
be rented out to people online, guests should be able to see what's available
, search for certain date they might want to stay with us, and
make booking and reservation
#### key functionality
* showcase the hotel
* allow for booking a room for one or more nights
* check a room's availability
* book the room
* notify guest, and notify owner when a reservation is successful (emails/text)
* have admin panel to the owner with authentication system
    * review existing bookings
    * show a calendar of bookings
    * change or cancel a booking

---
### HTTP endpoints

| HTTP Method | Resource                            | Description                                                                                                                                  |
|-------------|-------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------|
| GET         | /rooms                              | returns a list of rooms                                                                                                                      |
| POST        | /rooms                              | Creates a new room                                                                                                                           |
| GET         | /rooms/{room-id}                    | Returns a room by ID                                                                                                                         |
| POST        | /rooms/{room-id}/check-availability | Explode the url to extract room-id, Form that Submits start&end date<br/>and Response a Message feedback whether a room is available or not  |
|             |                                     |                                                                                                                                              | 
|             |                                     |                                                                                                                                              |
| PUT         | /admin/rooms/{room-id}              | Updates an existing room                                                                                                                     |
| DELETE      | /admin/rooms/{room-id}              | Deletes an existing room                                                                                                                     |
| GET         | /admin/reservations                 | returns a list of reservations                                                                                                               |
|             |                                     |                                                                                                                                              |
| GET         | /users/{user-id}/login              | Display the login page                                                                                                                       |
| POST        | /users/{user-id}/login              | Submit user login data                                                                                                                       |
| GET         | /users/{user-id}/logout             | Log user out                                                                                                                                 |
|             |                                     |                                                                                                                                              |
|             |                                     |                                                                                                                                              |
|             |                                     |                                                                                                                                              |
|             |                                     |                                                                                                                                              |



### Database Functions
| Name                                                       | Description                                      |
|------------------------------------------------------------|--------------------------------------------------|
| SelectRooms()([]*room, error)                              | Fetch all the rooms records                      |
| SelectRoom(id int) (*room, error)                          | Fetch a room by its id                           |
| InsertRoom() error                                         | Insert a new room                                |
| UpdateRoom(id int) error                                   | Update an existing room                          |
| DeleteRoom(id int) error                                   | Delete an existing room                          |
| SearchRoomAvailabilityByDatesByID(sd, ed, id)(*room,error) | Return count of match room to the sd, ed, and id |
| SearchRoomAvailabilityByID(id) (*room, error)              |                                                  |
|                                                            |                                                  |
|                                                            |                                                  |
|                                                            |                                                  |


`db function: get room by id from rooms table
for the returned room search in the roomRestrictions the start and end date
if a value returned then the room is reserved
if not then the room is free `
























