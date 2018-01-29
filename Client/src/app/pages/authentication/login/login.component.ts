import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, AbstractControl, Validators } from '@angular/forms';
import { AuthenticationService } from '../../../services/authentication.service';
import { UserCredential } from '../../../models/userCredentialModel';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  email: AbstractControl;
  password: AbstractControl;

  constructor(private fb: FormBuilder,
    private authService: AuthenticationService,
    private router: Router) { }

  ngOnInit() {
    this.loginForm = this.fb.group({
      'email': ['', Validators.compose([Validators.required, Validators.email])],
      'password': ['', Validators.compose([Validators.required])]
    })

    this.email = this.loginForm.controls['email'];
    this.password = this.loginForm.controls['password'];
  }

  getEmailErrorMessage() {
    return this.email.hasError('required') ? 'Must enter a value' :
      this.email.hasError('email') ? 'Not a valid format' : '';
  }

  getPasswordErrorMessage() {
    return this.password.hasError('required') ? 'Must enter a value' : '';
  }

  login() {
    if (!this.loginForm.valid) {
      return;
    }

    let user = new UserCredential();
    user.userName = this.email.value;
    user.password = this.password.value;

    this.authService.loginUser(user).subscribe((loginResponse) => {
      var loginResult = this.authService.processLoginResult(loginResponse);
      if (loginResult) {
        this.router.navigate(['chat']);
      }
    })
  }

  signup() {
    this.router.navigate(['auth', 'register']);
  }

}
