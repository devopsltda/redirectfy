import { CommonModule, DatePipe } from '@angular/common';
import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { CalendarModule } from 'primeng/calendar';
import { fadeInOutAnimation } from '../../animations/animations.module';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { MyValidatorsService } from '../../services/my-validators.service';

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [CommonModule,SharedModule,RouterModule,FormsModule,CalendarModule,ReactiveFormsModule,ToastModule],
  providers:[MessageService,DatePipe],
  animations:[fadeInOutAnimation],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.scss'
})
export class SignupComponent {
  formSubmited:boolean = false;
  formStep:string = 'init';
  signupForm!:FormGroup;
  formdata!:any
  hash:string = this.activatedRoute.snapshot.params['hashSignup']


  constructor(
    private router:Router,
    private formBuilder:FormBuilder ,
    private activatedRoute:ActivatedRoute,
    private redirectfyApi:RedirectifyApiService,
    private messageService:MessageService,
    private myValidators:MyValidatorsService
  )
    {

    this.signupForm = this.formBuilder.group(
      {
        data_de_nascimento:["",[Validators.required]],
        senha:["",[Validators.required,this.myValidators.minLength(8),this.myValidators.hasNumber(),this.myValidators.hasLetter()]],
        senha_confirmacao:["",[Validators.required]]
      }
    );
  }

    passwordsMatch():boolean {
    return this.signupForm.controls['senha'].value == this.signupForm.controls['senha_confirmacao'].value? true:false
  }

  async onSubmit(){
    this.formSubmited = true
      if(this.signupForm.valid && this.passwordsMatch()){
        try{
          this.formdata = this.signupForm.getRawValue()

          const apiResponse = await this.redirectfyApi.finishSignUp(this.hash,this.formdata['data_de_nascimento'],this.formdata['senha'])

          if(apiResponse.status == 200 ){
            this.router.navigate(['/signup/complete'])
          }

        } catch (error:any) {
          error.error.forEach(
            (errorMessage:string)=>{
            return this.messageService.add({severity:'error',summary:'Falha na Ação',detail:'Erro ao redefinir a senha'})
          }
        )

        }
      }
  }

}
