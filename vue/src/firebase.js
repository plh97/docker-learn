import {initializeApp} from 'firebase'


var app = initializeApp({
  apiKey: "AIzaSyCjWMja82tT5P1OpeMDhFGPSB2N-YsFGWg",
  authDomain: "pengliheng-company.firebaseapp.com",
  databaseURL: "https://pengliheng-company.firebaseio.com",
  projectId: "pengliheng-company",
  storageBucket: "pengliheng-company.appspot.com",
  messagingSenderId: "142728924372"
});


export const db = app.database();

export const namesRef = db.ref('names');


