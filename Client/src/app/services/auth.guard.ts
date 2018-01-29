import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { CurrentUser } from '../models/currentUserModel';
import { AuthenticationService } from './authentication.service';
import { Router } from '@angular/router';
import { CanActivateChild, CanLoad } from '@angular/router/src/interfaces';
import { Route } from '@angular/router/src/config';

@Injectable()
export class AuthGuard implements CanActivate, CanActivateChild, CanLoad {

  private loginUrl = ['auth', 'login']

  constructor(private authService: AuthenticationService, private router: Router) {
  }

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {
    if (this.authService.isLoggedIn()) { return true; }
    this.router.navigate(this.loginUrl, { queryParams: { redirectTo: state.url } });
  }

  canActivateChild(childRoute: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {
    return this.canActivate(childRoute, state)
  }

  canLoad(route: Route): Observable<boolean> | Promise<boolean> | boolean {
    if (this.authService.isLoggedIn()) {
      return true;
    }

    const url = `/${route.path}`;
    this.router.navigate(this.loginUrl, { queryParams: { redirectTo: url } });
    return false;
  }
}
