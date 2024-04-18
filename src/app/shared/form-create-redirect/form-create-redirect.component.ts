import { Component } from '@angular/core';
import { IconWhatsappComponent } from '../icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../icon-telegram/icon-telegram.component';
import { SharedModule } from '../shared.module';
import { AnimationsModule, fadeInOutAnimation } from '../../animations/animations.module';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-form-create-redirect',
  standalone: true,
  imports: [IconWhatsappComponent,IconTelegramComponent,SharedModule,AnimationsModule,ReactiveFormsModule,CommonModule,FormsModule],
  animations:[fadeInOutAnimation],
  templateUrl: './form-create-redirect.component.html',
  styleUrl: './form-create-redirect.component.scss'
})
export class FormCreateRedirectComponent {

  submitted:boolean = false

  formStep:string = 'init'
  getPlataforma!:string
  redirectName!:string
  prioridade:string = 'whatsapp,telegram'
  submitData:any = []
  createData:{[key:string]:any} = {}
  createLinkForm!:FormGroup

  constructor
  (
    private formBuilder:FormBuilder,
    private api:RedirectifyApiService,
    private router:Router

  )
  {
    this.createLinkForm = this.formBuilder.group({
      link:['',[Validators.required]],
      nome:[''],
      mensagem:[''],
      plataforma:[]
    })
  }

  isValid(){

  }
  createDataEmpty(){
    return Object.keys(this.createData).length?true:false
  }

  addContact(plataforma:string){
    this.submitted = true
    if(this.createLinkForm.valid){
      this.submitted = false
      if(plataforma == 'whatsapp'){
        const data = this.createLinkForm.getRawValue()
        data['plataforma'] = plataforma

        if(this.createData['whatsappData'] == undefined){
          this.createData['whatsappData'] = [data]
          this.createLinkForm.reset()
          this.formStep = 'init'
        } else{
          this.createData['whatsappData'].push(data)
          this.createLinkForm.reset()
          this.formStep = 'init'
        }

        console.log(this.createData['whatsappData'])
      } else if(plataforma == 'telegram'){

        const data = this.createLinkForm.getRawValue()

        data['plataforma'] = plataforma

        if(this.createData['telegramData'] == undefined){

          this.createData['telegramData'] = [data]
          this.createLinkForm.reset()
          this.formStep = 'init'
        } else{

          this.createData['telegramData'].push(data)
          this.createLinkForm.reset()
          this.formStep = 'init'
        }
        console.log(this.createData)
      }
    }
  }

  getContacts(plataforma:string){
    this.getPlataforma = plataforma
    this.formStep = 'getContacts'
  }

  generateRandomInteger(min: number, max: number): number {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

 async onSubmit(){
    for(let item of this.createData['whatsappData']){
      this.submitData.push(
        {
        nome:`${item.nome?item.nome:'+'+item.link}`,
        link:`https://wa.me/+${item.link}${item.mensagem ? `?text=${encodeURIComponent(item.mensagem)}` : ""}`,
        plataforma:item.plataforma
        }
      )
    }
    console.log(this.submitData)
    if(this.redirectName == undefined){
      this.redirectName = `Redirect #${this.generateRandomInteger(1,100)}`
    }

    try{
      const resApi = await this.api.createRedirect(this.redirectName,this.prioridade,this.submitData)
      if (resApi.status == 201) {
        this.router.navigate(['/home'])
      }
      console.log(resApi)
    } catch (error){
      console.log(error)
    }
  }

}
