# AIRSPACE DEMO API
### TIM LIEBERMAN   : )

#### OBJECTIVE: 
Using N_S_UAS_Flight_Restriction data create a api that takes a lat long and responds with message revealing if cords
are in restricted airspace or not. 

### PRE-Reqs
#### Required To Run:
I'm presuming you already have these tools installed but if not i've provided some psuedo installs.  They might be accurate.

- `brew install docker`
- `brew install go`

#### Testing: 
The only prerequisite for testing is that you have gomock installed... 
- ```go install github.com/golang/mock/mockgen@latest```

### LOCAL-Build

1. From project root:
   2. `docker compose up --build`
      3. No Errors:
         4. I wasn't having luck loading in the data via script.  As of right now you'll need to run:
            5. `docker ps` grab the postgres container id.
            6. `docker exec -it bash` opens up the shell
            7. `./load_geojson.sh`
