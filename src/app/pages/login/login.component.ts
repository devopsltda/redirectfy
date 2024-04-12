import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { InputTextModule } from 'primeng/inputtext';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { MessageService } from 'primeng/api';
import { error } from 'node:console';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [SharedModule,InputTextModule,CommonModule,RouterLink, ReactiveFormsModule],
  providers:[MessageService],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent  {
  loginForm!:FormGroup
  loginData!:any;
  formSubmited:boolean = false;

  constructor( private formBuilder:FormBuilder , private redirectifyApi:RedirectifyApiService, private messageService:MessageService ){
    this.loginForm = this.formBuilder.group({
      email:['pabloed0009@gmail.com',[Validators.required,Validators.email]],
      senha:['teste123',[Validators.required]]
    })
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
        console.log(reqLogin.headers.getAll('Set-Cookie'))

      } catch(error:any){
        if(error.status == 400){
          console.log(error)
          this.showError('Erro no Login','Email ou Senha incorretos')
        }
      }
    }

  }
}
