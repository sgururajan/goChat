import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, AbstractControl, Validators } from '@angular/forms';
import { matchValidator } from '../../../validators/matchValidator';
import { AuthenticationService } from '../../../services/authentication.service';
import { RegisterUser } from '../../../models/userRegisterModel';
import { LoginResponse } from '../../../models/loginResponseModel';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

  registerForm: FormGroup;
  email: AbstractControl;
  password: AbstractControl;
  confirmPassword: AbstractControl;
  firstName: AbstractControl;
  lastName: AbstractControl;
  nickName: AbstractControl;

  constructor(private fb: FormBuilder,
    private authService: AuthenticationService,
    private router: Router) { }

  ngOnInit() {
    this.registerForm = this.fb.group({
      'email': ['', Validators.compose([Validators.required, Validators.email])],
      'password': ['', Validators.compose([Validators.required])],
      'confirmPassword': ['', Validators.compose([Validators.required, matchValidator('password')])],
      'firstName': ['', Validators.compose([Validators.required])],
      'lastName': [''],
      'nickName': ['']
    })

    this.email = this.registerForm.controls['email']
    this.password = this.registerForm.controls['password']
    this.confirmPassword = this.registerForm.controls['confirmPassword']
    this.firstName = this.registerForm.controls['firstName'];
    this.lastName = this.registerForm.controls['lastName'];
    this.nickName = this.registerForm.controls['nickName']
  }

  getEmailErrorMessage() {
    return this.email.hasError('required') ? 'You must enter a value' :
      this.email.hasError('email') ? 'Not a valid format' : ''
  }

  getConfirmPasswordErroMessage() {
    return this.confirmPassword.hasError('required') ? 'You must enter a value' :
      this.confirmPassword.hasError('notEqual') ? 'Not a match' : ''
  }

  getFirstNameErrorMessage() {
    return this.firstName.hasError('required') ? 'You must enter a value' : ''
  }

  getPasswordErrorMessage() {
    return this.password.hasError('required') ? 'You must enter a value' : ''
  }

  onSubmit() {
    if (!this.registerForm.valid) { return; }
    let user = new RegisterUser()
    user.email = this.email.value;
    user.passwordHashed = this.password.value;
    user.firstName = this.firstName.value;
    user.lastName = this.lastName.value;
    user.nickName = this.nickName.value;

    this.authService.registerUser(user).subscribe((o) => {
      const loginResult = this.authService.processLoginResult(o);
      if (loginResult) {
        this.router.navigate(['chat']);
      }
    });
  }

  login() {
    this.router.navigate(['auth', 'login'])
  }
}
