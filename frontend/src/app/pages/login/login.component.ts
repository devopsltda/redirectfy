import { Component, OnInit } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { InputTextModule } from 'primeng/inputtext';
import { CommonModule } from '@angular/common';
import { Router, RouterLink } from '@angular/router';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { MessageService } from 'primeng/api';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [SharedModule,InputTextModule,CommonModule,RouterLink, ReactiveFormsModule],
  providers:[MessageService],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent implements OnInit {
  loginForm!:FormGroup
  loginData!:any;
  formSubmited:boolean = false;
  isAuthenticated:boolean = this.cookie.check('refresh-token') && this.cookie.check('access-token') && this.cookie.check('usuario')
  constructor
  (
    private formBuilder:FormBuilder ,
    private redirectifyApi:RedirectifyApiService,
    private messageService:MessageService,
    private cookie:CookieService,
    private router:Router
  )
    {
    this.loginForm = this.formBuilder.group({
      email:['',[Validators.required,Validators.email]],
      senha:['',[Validators.required]]
    })
  }

  ngOnInit(): void {
    if(this.isAuthenticated){
      this.router.navigate(['/home'])
    }
  }

  invalidTouchedDirtyControl(controlador:string){
    return this.loginForm.controls[controlador].invalid && (this.loginForm.controls[controlador].touched || this.loginForm.controls[controlador].dirty);
  }

  showError(title:string,message:string){
    this.messageService.add({severity: 'error', summary: title, detail: message});
  }

async onSubmit(){
    this.formSubmited = true
    if(this.loginForm.valid){
    try{
        this.loginData = this.loginForm.getRawValue()
        const reqLogin = await this.redirectifyApi.login(this.loginData['email'],this.loginData['senha'])
        if(reqLogin.status == 200)
          {
          this.router.navigate(['/home'])
          }
      } catch(error:any){
        if(error.status == 400){
          this.showError('Login Invalido','Email ou Senha incorretos')
        }
        this.showError('Login Invalido','Falha ao realizar o login')
      }
    }

  }
}
