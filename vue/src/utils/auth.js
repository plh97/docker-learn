import {auth} from 'firebase'
var uiConfig = {
  signInSuccessUrl: '/',
  signInOptions: [
    // Leave the lines as is for the providers you want to offer your users.
    auth.GoogleAuthProvider.PROVIDER_ID,
    auth.GithubAuthProvider.PROVIDER_ID,
  ],
  // tosUrl and privacyPolicyUrl accept either url string or a callback
  // function.
  // Terms of service url/callback.
  tosUrl: '/',
  // Privacy policy url/callback.
  privacyPolicyUrl: function() {
    window.location.assign('/');
  }
};
export default uiConfig

