import {initializeApp} from 'firebase'


var app = initializeApp({
  apiKey: "AIzaSyBBAclYHhcMIaXu_Avfrxsf5Dkud5xrDc8",
  authDomain: "vue-company-849f5.firebaseapp.com",
  databaseURL: "https://vue-company-849f5.firebaseio.com",
  projectId: "vue-company-849f5",
  storageBucket: "",
  messagingSenderId: "306590538429"
});


export const db = app.database();

export const namesRef = db.ref('names');


