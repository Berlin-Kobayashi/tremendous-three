# tremendous-three


Variables:

No. of rides integer
No. of vehicles integer

City 2 dimensional array

vehicle to rides one to many relationship

Each ride:
    distance int
    start_intersection obj/arr of 2
    finish_intersection obj/arr of 2
    start_actual step int
    finish_actual step int
    start_earliest step int
    finish_latest step int
    
Conditions mandatory:
    start_actual >= start_earliest
    
Conditions efficient:
    finish_actual <= finish_latest
    
    
For every ride that finishes on time (or early), you will earn points proportional to the distance of that ride;
plus an additional bonus if the ride also started precisely on time.