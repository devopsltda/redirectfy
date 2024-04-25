import { Component } from '@angular/core';
import { IconWhatsappComponent } from '../icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../icon-telegram/icon-telegram.component';
import { SharedModule } from '../shared.module';
import { AnimationsModule, fadeInOutAnimation } from '../../animations/animations.module';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { Router, RouterLink } from '@angular/router';
import { MessageService } from 'primeng/api';


@Component({
  selector: 'app-form-create-redirect',
  standalone: true,
  imports: [IconWhatsappComponent,IconTelegramComponent,SharedModule,AnimationsModule,ReactiveFormsModule,CommonModule,FormsModule,RouterLink],
  providers:[MessageService],
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
  whatsappForm!:FormGroup
  telegramForm!:FormGroup

  constructor
  (
    private formBuilder:FormBuilder,
    private api:RedirectifyApiService,
    private router:Router,
    private messageService:MessageService
  )
  {
    this.whatsappForm = this.formBuilder.group({
      link:['',[Validators.required,Validators.pattern(/^\d{13,}$/)]],
      nome:['',[Validators.required]],
      mensagem:[''],
      plataforma:[]
    })

    this.telegramForm = this.formBuilder.group({
    link:['',[Validators.required,Validators.pattern(/^https:\/\/t\.me\/.*/ )]],
      nome:['',[Validators.required]],
      plataforma:[]
    })

  }

  createDataEmpty(){
    return Object.keys(this.createData).length?true:false
  }
  arrayLength(data:any[]){
    return data.length
  }

  addContact(plataforma:string){
    this.submitted = true

    if(this.whatsappForm.valid || this.telegramForm.valid){
      this.submitted = false
      if(plataforma == 'whatsapp'){
        const data = this.whatsappForm.getRawValue()
        data['plataforma'] = plataforma
        if(this.createData['whatsappData'] == undefined){
          this.createData['whatsappData'] = [data]
          this.whatsappForm.reset()
          this.formStep = 'init'
        } else{
          this.createData['whatsappData'].push(data)
          this.whatsappForm.reset()
          this.formStep = 'init'
        }
      }



      else if(plataforma == 'telegram'){
        const data = this.telegramForm.getRawValue()
        data['plataforma'] = plataforma
        if(this.createData['telegramData'] == undefined){

          this.createData['telegramData'] = [data]
          this.telegramForm.reset()
          this.formStep = 'init'
        } else{
          this.createData['telegramData'].push(data)
          this.telegramForm.reset()
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
  if(this.createData['whatsappData']){
    for(let item of this.createData['whatsappData']){
      this.submitData.push(
        {
        nome:`${item.nome?item.nome:'+'+item.link}`,
        link:`https://wa.me/+${item.link}${item.mensagem ? `?text=${encodeURIComponent(item.mensagem)}` : ""}`,
        plataforma:item.plataforma
        }
      )
    }
  }
    if(this.createData['telegramData']){
      for(let item of this.createData['telegramData']){
        this.submitData.push(
          {
          nome:`${item.nome?item.nome:item.link}`,
          link:item.link,
          plataforma:item.plataforma
          }
        )
      }
    }

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
      this.messageService.add({summary:"Falha ao Criar Redirecionador",detail:'Ocorreu um erro ao criar o redirecionador, ação não executada',severity:'error'})
    }
  }

}
