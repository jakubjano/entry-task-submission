# Entry task submission

## Event aggregator

Here is my solution to the assigned entry task. I have tried to replicate the example output \
from the assignment without any fancy additions. I tried to structure the application according to \
known concepts in backend development.
```
|-Project
    |-endpoints
        |-events
        |-helpers
    |-logic
        |-models
    main
```
Events with timestamps of their arrival are saved at the `POST` endpoint\
and aggregated data, based on the requested query, are returned in a response\
at the `GET` endpoint. Communication with the service is handled by the corresponding\
handlers and logic of the aggregation is handled in a separate file.

Basic validation function is called on the query parameters before aggregation.

Unfortunately I have not been able to tackle the bonus part, since I have little to none experience\
with testing in GO and found it rather difficult. 




