
import * as functions from 'firebase-functions';
import * as dik from './dikshaneree';

export const werd = (s:string) => {
  const result = dik.dikshaneree[s.toUpperCase()] || s;
  return result;
}

export const reespel = functions.https.onRequest((request, response) => {
    response.send(werd("Hello") + " " + werd("from") + " "+ werd("Firebase") + "!");
});
