# Assignment1

## Description
REST web application that uses two APis to provide a client that retrieves information about universities that may be candidates for applicatin based on the country it is located in.

***APIs used:***
- https://restcountries.eu/
  -    https://github.com/Hipo/university-domains-list/
- https://universities.hipolabs.com/
  -  https://gitlab.com/amatos/rest-countries

## Endpoints

```
http://localhost:8080/unisearcher/v1/uniinfo/
http://localhost:8080/unisearcher/v1/neighbourunis/
http://localhost:8080/unisearcher/v1/diag/
```

## Deployment
- clone the repository
- open in the IDE your choice and run /cmd/main.go
  - If you are not using an integrated development environment (IDE), follow these steps to compile the source code:

    1.  Navigate to the source directory.
    
    2. Run the following command for Linux: "go build -o diag .". For Windows: "go build -o diag.exe .".
    3. If you are using Linux, make sure to add execution permissions using the command "chmod 740 diag".
    
    4. After compiling the source code, you can run the executable binary using the following commands: 
       - For Linux: "./diag".
       - For Windows: "diag.exe".

## Body

### Example response

```` json
        "name": "Universiti Teknikal Malaysia Melaka",
        "country": "Malaysia",
        "isocode": "MY",
        "webpages": [
            "http://www.utem.edu.my/"
        ],
        "languages": {
            "eng": "English",
            "msa": "Malay"
        },
        "map": {
            "googleMaps": "https://goo.gl/maps/qrY1PNeUXGyXDcPy6",
            "openStreetMaps": "https://www.openstreetmap.org/relation/2108121"
        }
````

## Example request

```
http://localhost:8080/unisearcher/v1/uniinfo/tek
http://localhost:8080/unisearcher/v1/neighbourunis/norway/tech?limit=5
http://localhost:8080/unisearcher/v1/diag/
```

## Authors
**Magnus W. Johannessen**

## Credits
**Christopher Frantz**