# Medicare

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/muskankhedia/Revolutionize_Medicare/graphs/commit-activity) [![GitHub license](https://img.shields.io/github/license/muskankhedia/Revolutionize_Medicare.svg)](https://github.com/muskankhedia/Revolutionize_Medicare/blob/master/LICENSE) [![GitHub contributors](https://img.shields.io/github/contributors/muskankhedia/Revolutionize_Medicare.svg)](https://GitHub.com/muskankhedia/Revolutionize_Medicare/graphs/contributors/)


[![GitHub stars](https://img.shields.io/github/stars/muskankhedia/Revolutionize_Medicare.svg?style=social&label=Star&maxAge=2592000)](https://GitHub.com/muskankhedia/Revolutionize_Medicare/stargazers/) [![GitHub forks](https://img.shields.io/github/forks/muskankhedia/Revolutionize_Medicare.svg?style=social&label=Fork&maxAge=2592000)](https://GitHub.com/muskankhedia/Revolutionize_Medicare/network/) [![GitHub watchers](https://img.shields.io/github/watchers/muskankhedia/Revolutionize_Medicare.svg?style=social&label=Watch&maxAge=2592000)](https://GitHub.com/muskankhedia/Revolutionize_Medicare/watchers/)

Medicare is an application that aims to modernize and revolutionize the way that Healthcare works in India. It reduces the information gap between a patient and a doctor, **significantly reducing the chances of wrong diagnosis**. Medicare helps to ease the complex medical decisions and calculations with the help of deep learning and secure the entire process by implementing it in a blockchain. Our model considers information like the patient's record of previous ailments, his Sugar level, Blood Pressure, BMI, TPR, etc to smartly suggest medicines and predict the result of taking said medication. The crucial patient records (i.e., his medical history) is stored using a blockchain in order to be secure, scalable and fault tolerant.

This project aspires to bring the best of the smart world to people of all strata. Our software is unique in the fact that people be it rich or poor will be equally benefited. This would bring about a revolution in the way people look for better treatments. More than 7000 deaths are caused in India every year because someone couldn’t read a Doctor’s prescription or the Doctor prescribed a medicine that was unsuitable. Most people in general aren’t very sincere when it comes to maintaining their medical records, which can greatly influence the medication that is suitable for them. For example, a medicine may not work on a patient with a history of blood sugar but may work on others. Medicare will fully utilize this precious data. Its digital blockchain library of patient’s medical history and deep learning based model to predict the effectiveness and suggest medicines will hugely influence the way Doctor’s treat patients and prescribe medicines.

## :minidisc: Installation instructions
You must have [Go](https://golang.org/) and [npm](https://www.npmjs.com/) installed in your computer. Then follow these steps:

```
go get github.com/muskankhedia/Revolutionize_Medicare
```

Install all Go dependencies by running
```
go get github.com/patrikeh/go-deep
go get github.com/patrikeh/go-deep/training
go get github.com/gorilla/mux
```
Install `http-server` by running the command
```
npm install -g http-server
```

Run your http-server by executing the command 
```
npm start
```
and your Go server by running the following command in your `service` folder
```
go run main.go
```



## :tada: How it works
The following patient data is stored.

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

>>Medicare helps the doctor prescribe medicines to a patient based on this medical history. The doctor can also judge the effectiveness of a particular medicine on similar patients  based on its history of success/ failure.

**The effectiveness of Medicines are highly dependent on factors like sugar, BP, TPR which are often the cause of failure of medicine. Hence, different medicines are prescribed for patients suffering with these. Choosing the right medicine for these complicated cases is tough for human minds, hence we propose to solve this crucial issue by the help of `deep learning`**.

## :wrench: Technology Stack
* **Backend** Go
* **Front-end** AngularJS
* **Database** JSON-db (in-memory)
* `Deep Learning` model used to predict the success rate of a medicine with respect to a particular disease.
* `Block Chain` is used to store the medical history of each Patient.
* Hosted using virtual machines deployed at `AWS cloud` at [link](http://54.237.215.120:8080).

## :rocket: Future Scope

### Things to be implemented
 - [ ] Epidimic prediction based on the number of cases of a particular disease in an area.
 - [ ] Complete support for smart medicine suggestion considering `allergies`.
 - [ ] Graphical view of user Medical History.
 - [ ] Feedback loop on diagnosis, thus validating both the Doctor and the medication.

## AWS Services and how it helped us
AWS has been the perfect solution for our project. It helped us deploy our instances on virtual machines and was then linked with AWS simple storage bucket for managing our files. We had used aws hosted mongodb at the initial stage of development to carry out the tests and compatibilities on our devices. AWS provided us the needed scalability at a much economical cost, in a user-friendly way, which was easier to adopt and grow with our busy time in the hack. AWS’s zero downtime (almost), fairer billing services, high scalability makes it a fit cloud provider.

## Tracks used
- [ ] Revolutionize The Smart World

## Demo video
Click on this [link](https://drive.google.com/file/d/11G2IIeCx2ac1Hp1GGssfKoTBXU-2CZrh/view) to watch a demo on the application.

## :gem: Contributors
Developed with :hearts: by team CodeZero.
1. [Muskan Khedia](https://github.com/muskankhedia) - Core Developer, Maintainer
2. [Harkishen Singh](https://github.com/Harkishen-Singh) - Core Developer, Maintainer
3. [C. Anirudh](https://github.com/C-Anirudh) - Developer