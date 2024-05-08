import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { MessageService } from 'primeng/api';
import { MyValidatorsService } from '../../services/my-validators.service';


@Component({
  selector: 'app-reset-password',
  standalone: true,
  imports: [SharedModule,CommonModule,RouterModule,FormsModule,ReactiveFormsModule],
  providers:[MessageService],
  templateUrl: './reset-password.component.html',
  styleUrl: './reset-password.component.scss'
})
export class ResetPasswordComponent {

  formSubmited:boolean = false;
  formStep:string = 'init';
  resetPassForm!:FormGroup;
  formdata!:any
  hash:string = this.activatedRoute.snapshot.params['hash']

  constructor(
    private router:Router,
    private formBuilder:FormBuilder ,
    private activatedRoute:ActivatedRoute,
    private redirectfyApi:RedirectifyApiService,
    private messageService:MessageService,
    private myValidators:MyValidatorsService){

      this.resetPassForm = this.formBuilder.group(
        {
          senha:["",[Validators.required,this.myValidators.minLength(8),this.myValidators.hasNumber(),this.myValidators.hasLetter()]],
          senha_confirmacao:["",[Validators.required]]
        }
      );
  }


    passwordsMatch():boolean {
    return this.resetPassForm.controls['senha'].value == this.resetPassForm.controls['senha_confirmacao'].value? true:false
  }

  async onSubmit(){
    this.formSubmited = true
      if(this.resetPassForm.valid && this.passwordsMatch()){
        try{
          this.formdata = this.resetPassForm.getRawValue()

          const apiResponse = await this.redirectfyApi.resetPassword(this.hash,this.formdata['senha'])

          if(apiResponse.status == 200 ){
            this.router.navigate(['/signup/complete'])
          }

        } catch (error:any) {
          error.error.forEach(
            (errorMessage:string)=>{
            return this.messageService.add({severity:'error',summary:'Falha na Ação',detail:'Erro ao finalizar o cadastro'})
          }
        )

        }
      }
  }

}
