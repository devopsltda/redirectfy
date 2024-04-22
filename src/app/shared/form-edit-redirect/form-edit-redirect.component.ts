import { Component, OnInit } from '@angular/core';
import { IconWhatsappComponent } from '../icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../icon-telegram/icon-telegram.component';
import { SharedModule } from '../shared.module';

import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { fadeInOutAnimation } from '../../animations/animations.module';
import { Hash } from 'crypto';
import { get } from 'http';
import { MessageService } from 'primeng/api';


@Component({
  selector: 'app-form-edit-redirect',
  standalone: true,
  imports: [IconWhatsappComponent,IconTelegramComponent,SharedModule,ReactiveFormsModule,CommonModule,FormsModule,RouterLink],
  providers:[MessageService],
  animations:[fadeInOutAnimation],
  templateUrl: './form-edit-redirect.component.html',
  styleUrl: './form-edit-redirect.component.scss'
})
export class FormEditRedirectComponent implements OnInit {
  editedLinks = []
  redirectHash:string = this.activatedRoute.snapshot.params['hash_redirect']

  disableEditNome:boolean = true;
  submitted:boolean = false

  redirectName!:string
  formStep:string = 'init'
  getPlataforma!:string

  prioridade:string = 'whatsapp,telegram'

  submitData:any = []
  editedData:any = {'whatsappData':[],'telegramData':[]}
  redirectData!:any

  createData:{[key:string]:any} = {}
  whatsappForm!:FormGroup
  telegramForm!:FormGroup



  constructor

  (
    private formBuilder:FormBuilder,
    private api:RedirectifyApiService,
    private router:Router,
    private messageService:MessageService,
    private activatedRoute:ActivatedRoute
  )

  {
    this.whatsappForm = this.formBuilder.group({
      numero:['',[Validators.required,Validators.pattern(/^\d{13,}$/)]],
      nome:['',[Validators.required]],
      mensagem:[''],
      plataforma:[],
      id:[]
    })

    this.telegramForm = this.formBuilder.group({
    link:['',[Validators.required,Validators.pattern(/^https:\/\/t\.me\/.*/ )]],
      nome:['',[Validators.required]],
      plataforma:[],
      id:[]
    })
  }

 async ngOnInit() {
    await this.getRedirectData();
    this.redirectName = this.redirectData.redirecionador.nome
    console.log(this.redirectData)

  }

  objectLength(data:any){
    return Object.values(data).length
  }

  async getRedirectData(){
    this.redirectData = await this.api.getRedirect(this.redirectHash)
    const whatsappData = this.redirectData.links.filter((link:any) => link.plataforma ==='whatsapp')
    const telegramData = this.redirectData.links.filter((link:any) => link.plataforma ==='telegram')
    Object.assign(this.redirectData,{whatsappData,telegramData});
    delete this.redirectData.links

    return  this.redirectData
  }


  createDataEmpty(){
    return Object.keys(this.createData).length?true:false
  }


  getContacts(plataforma:string){
    this.getPlataforma = plataforma
    this.formStep = 'getContacts'
  }

  generateRandomInteger(min: number, max: number): number {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  redirectDataEmpty(){
    return Object.keys(this.redirectData).length?true:false
  }

 async buttonCardEvent(event:string,data:any){
    console.log(event)
    if(event == 'editar'){
      if(data.plataforma == 'whatsapp'){
        const numero = data.link.match(/\/\+(\d+)/)[1];
        let mensagem = data.link.match(/text=([^&]*)/)
        if (mensagem) {
          mensagem = decodeURIComponent(mensagem[1]);
      }
        this.whatsappForm.controls['numero'].setValue(numero);
        this.whatsappForm.controls['mensagem'].setValue(mensagem);
        this.whatsappForm.controls['nome'].setValue(data.nome)
        this.whatsappForm.controls['id'].setValue(data.id)
        this.formStep = 'editW'
      }
      else if(data.plataforma == 'telegram'){
       this.telegramForm.controls['link'].setValue(data.link);
       this.telegramForm.controls['nome'].setValue(data.nome)
       this.telegramForm.controls['id'].setValue(data.id)
       this.formStep = 'editT'
     }
    }
    else if(event == 'enable'){
      try {
        await this.api.enableLinkInRedirect(this.redirectHash,data.id)
        await this.ngOnInit()
      } catch (error) {
        console.log(error)
      }
    }
    else if(event == 'disable'){
      try {
        await this.api.disableLinkInRedirect(this.redirectHash,data.id)
        await this.ngOnInit()
      } catch (error) {
        console.log(error)
      }
    }
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


  async saveEdits(plataforma:string){
    this.submitted = true
    if(plataforma == 'whatsapp'){
      if(this.whatsappForm.valid){
        let editedData = this.whatsappForm.getRawValue()
        let submitEditData
        for(let item of  Object.keys(editedData)){
          const item = editedData
          submitEditData =
            {
              nome:`${item.nome?item.nome:'+'+item.link}`,
              link:`https://wa.me/+${item.numero}${item.mensagem ? `?text=${encodeURIComponent(item.mensagem)}` : ""}`,
              plataforma:item.plataforma
            }
        }
        try{
          await this.api.updateLinkInRedirect(this.redirectHash,editedData['id'],submitEditData)
          await this.ngOnInit()
        } catch (error){
          console.log(error)
        }
        this.formStep = 'init'
        this.whatsappForm.reset()
      }
    } else  if(plataforma == 'telegram'){
      if(this.telegramForm.valid){
        let editedData = this.telegramForm.getRawValue()
        console.log(editedData)
        try{
          await this.api.updateLinkInRedirect(this.redirectHash,editedData['id'],editedData)
          await this.ngOnInit()
        } catch (error){
          console.log(error)
        }
        this.telegramForm.reset()
        this.formStep = 'init'
      }
    }
  }

 async onSubmit(){
  if(this.createData['whatsappData']){
    for(let item of this.createData['whatsappData']){
      this.submitData.push(
        {
        nome:`${item.nome?item.nome:'+'+item.numero}`,
        link:`https://wa.me/+${item.numero}${item.mensagem ? `?text=${encodeURIComponent(item.mensagem)}` : ""}`,
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
    try{
      const resApi = await this.api.addLinkToRedirect(this.redirectHash,this.submitData)
      if (resApi.status == 201) {
        await this.ngOnInit()
        this.formStep = 'init'
      }
    } catch (error){
      this.messageService.add({summary:"Falha ao Criar Redirecionador",detail:'Ocorreu um erro ao criar o redirecionador, ação não executada',severity:'error'})
    }
  }

}
