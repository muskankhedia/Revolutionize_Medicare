# Medicare

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/muskankhedia/Revolutionize_Medicare/graphs/commit-activity) [![GitHub license](https://img.shields.io/github/license/muskankhedia/Revolutionize_Medicare.svg)](https://github.com/muskankhedia/Revolutionize_Medicare/blob/master/LICENSE) [![GitHub contributors](https://img.shields.io/github/contributors/muskankhedia/Revolutionize_Medicare.svg)](https://GitHub.com/muskankhedia/Revolutionize_Medicare/graphs/contributors/)


[![GitHub stars](https://img.shields.io/github/stars/muskankhedia/Revolutionize_Medicare.svg?style=social&label=Star&maxAge=2592000)](https://GitHub.com/muskankhedia/Revolutionize_Medicare/stargazers/) [![GitHub forks](https://img.shields.io/github/forks/muskankhedia/Revolutionize_Medicare.svg?style=social&label=Fork&maxAge=2592000)](https://GitHub.com/muskankhedia/Revolutionize_Medicare/network/) [![GitHub watchers](https://img.shields.io/github/watchers/muskankhedia/Revolutionize_Medicare.svg?style=social&label=Watch&maxAge=2592000)](https://GitHub.com/muskankhedia/Revolutionize_Medicare/watchers/)

Medicare is an application that aims to modernize and revolutionize the way that Healthcare works in India. It reduces the information gap between a patient and a doctor, reducing the chances of wrong diagnosis significantly. It helps to simplify the complex task of choosing medication and other patient related decisions with the help of Deep learning. The crucial patient records (i.e., his medical history) is stored using a blockchain.

## :disc: Installation instructions
You must have [Go]() and [npm]() installed in your computer. Then follow these steps:

```
go get github.com/muskankhedia/Revolutionize_Medicare
```

Install all Go dependencies by running
```
go get github.com/patrikeh/go-deep
go get github.com/patrikeh/go-deep/training
go get github.com/gorilla/mux
```




## :tada: How it works
Entire medical history of a patient is stored (entered by the doctor every time he visits one). For a newborn child, complete details starting from vaccinations taken will be stored.

- [ ] The basic medical data stored includes :
   - [x] Basic Data
   - [x] Age
   - [x] Gender
   - [x] Sugar Level
   - [x] Blood Pressure
   - [x] BMI
   - [x] TPR (Temperature, Pulse and Respiration rate)
   - [x] Allergies
- [ ] Disease data (entered by doctor)
   - [x] Disease name
   - [x] Medication prescribed
   - [x] Duration of the ailment
   - [x] Boolean indicating whether medication was successful  

>>Medicare helps the doctor prescribe medicines to a patient based on this medical history (previous diseases and ailments) and their current bodily condition (that includes blood sugar level, blood pressure, etc).
The doctor can also judge the effectiveness of a particular medicine on similar patients  based on its history of success/ failure.

**The effectiveness of Medicines are highly dependent on factors like sugar, BP, TPR which are often the cause of failure of medicine. Hence, different medicines are prescribed for patients suffering with these. Choosing the right medicine for these complicated cases is tough for human minds, hence we propose to solve this crucial issue by the help of `deep learning`**.

## :wrench: Technology Stack
* **Backend** Go
* **Front-end** AngularJS
* `Deep Learning` model used to predict the success rate of a medicine with respect to a particular disease.
* A `Block Chain` is used to store the medical history of each Patient.

## :gem: Contributors

1. [Muskan Khedia](https://github.com/muskankhedia) - Core Developer, Maintainer
2. [Harkishen Singh](https://github.com/Harkishen-Singh) - Core Developer, Maintainer
3. [C. Anirudh](https://github.com/C-Anirudh) - Developer