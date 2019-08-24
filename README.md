# Medicare

Medicare is an application that aims to modernize and revolutionize the way that Healthcare works in India. It reduces the information gap between a patient and a doctor, reducing the chances of wrong diagnosis significantly.

## :wrench: Technology Stack
* **Backend** Go
* **Front-end** AngularJS

## :tada: How it works
Entire medical history of a patient is stored (entered by the doctor every time he visits one). For a newborn child, complete details starting from vaccinations taken will be stored.

[] The medical data stored includes :
        [x] Basic Data
        [x] Age
        [x] Gender
        [x] Sugar Level
        [x] Blood Pressure
        [x] BMI
        [x] TPR (Temperature, Pulse and Respiration rate)
        [x] Allergies
[] Disease data (entered by doctor)
        [x] Disease name
        [x] Medication prescribed
        [x] Duration of the ailment
        [x] Boolean indicating whether medication was successful  

>>Medicare helps the doctor prescribe medicines to a patient based on this medical history (previous diseases and ailments) and their current bodily condition (that includes blood sugar level, blood pressure, etc).
>>The doctor can also judge the effectiveness of a particular medicine on similar patients  based on its history of success/ failure.
The effectiveness of Medicines are highly dependent on factors like sugar, BP, TPR which are often the cause of failure of medicine. Hence, different medicines are prescribed for patients suffering with these. Choosing the right medicine for these complicated cases is tough for human minds, hence we propose to solve this crucial issue by the help of `deep learning`.
