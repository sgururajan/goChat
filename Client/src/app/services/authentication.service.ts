import { Injectable } from '@angular/core';
import { RegisterUser } from '../models/userRegisterModel';
import { Observable } from 'rxjs/Observable';
import { HttpClient } from '@angular/common/http';
import { LoginResponse } from '../models/loginResponseModel';
import { CurrentUser } from '../models/currentUserModel';
import { Router } from '@angular/router';
import { UserCredential } from '../models/userCredentialModel';

@Injectable()
export class AuthenticationService {

  private registerApi = '/api/register';
  private loginApi = '/api/login';

  constructor(
    private httpClient: HttpClient,
    private router: Router) { }

  registerUser(user: RegisterUser) {
    return this.httpClient.post<LoginResponse>(this.registerApi, user);
  }

  loginUser(user: UserCredential) {
    return this.httpClient.post<LoginResponse>(this.loginApi, user);
  }

  processLoginResult(loginResponse: LoginResponse): Observable<boolean> | Promise<boolean> | boolean {
    if (!loginResponse.accessToken || loginResponse.accessToken.length == 0) {
      return false;
    }

    let currentUser = loginResponse as CurrentUser;
    localStorage.setItem('currentUser', JSON.stringify(currentUser));
    localStorage.setItem('access_token', currentUser.accessToken);
    return true;
  }

  isLoggedIn(): Observable<boolean> | Promise<boolean> | boolean {
    var token = localStorage.getItem('access_token');
    return token && token.length > 0;
    // var currentUser = JSON.parse(localStorage.getItem('currentUser')) as CurrentUser;
    // return currentUser && currentUser.accessToken.length > 0;
  }

  logout() {
    localStorage.removeItem('access_token');
    localStorage.removeItem('currentUser');
    return true;
  }

}
