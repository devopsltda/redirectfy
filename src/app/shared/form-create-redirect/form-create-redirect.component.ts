import { Component } from '@angular/core';
import { IconWhatsappComponent } from '../icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../icon-telegram/icon-telegram.component';
import { SharedModule } from '../shared.module';
import { AnimationsModule, fadeInOutAnimation } from '../../animations/animations.module';
import { RouterModule } from '@angular/router';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { isEmpty } from 'rxjs';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-form-create-redirect',
  standalone: true,
  imports: [IconWhatsappComponent,IconTelegramComponent,SharedModule,AnimationsModule,RouterModule,ReactiveFormsModule,CommonModule],
  animations:[fadeInOutAnimation],
  templateUrl: './form-create-redirect.component.html',
  styleUrl: './form-create-redirect.component.scss'
})
export class FormCreateRedirectComponent {
  formStep:string = 'init'
  createData:{[key:string]:any} = {}
  createLinkForm!:FormGroup
  getPlataforma!:string

  constructor
  (
    private formBuilder:FormBuilder
  )
  {
    this.createLinkForm = formBuilder.group({
      link:[],
      nome:[],
      plataforma:[]
    })
  }

  createDataEmpty(){
    return Object.keys(this.createData).length?true:false
  }

  addContact(plataforma:string){


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

      console.log(this.createData)
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

  getContacts(plataforma:string){
    this.getPlataforma = plataforma
    this.formStep = 'getContacts'

  }

}
