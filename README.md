# Workdaze-api

A public API that returns workdays for a given month

Example of endpoint:
https://workdaze-api-j5fajyv6ua-uk.a.run.app/api?year=2022&month=06&holidays=NewYear

API format:
 - year=YYYY
 - month = MM
 - holidays = 
   List of holidays:
    - NewYear"
    - MLK
    - Presidents
    - Memorial
    - Juneteenth
    - Indenpendence `use this misspelling`
    - Labor
    - Columbus
    - Veterans
    - Thanksgiving
    - Christmas
    
  Returns:
  ```{
    "Year": "2022",
    "month": "06",
    "holidays": [
        "NewYear",
        "MLK",
        "Presidents",
        "Memorial",
        "Juneteenth",
        "Indenpendence",
        "Labor",
        "Columbus",
        "Veterans",
        "Thanksgiving",
        "Christmas"
    ],
    "Days": 21
}
