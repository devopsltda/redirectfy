import { Component, OnInit } from '@angular/core';
import { IconWhatsappComponent } from '../icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../icon-telegram/icon-telegram.component';
import { SharedModule } from '../shared.module';

import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { ActivatedRoute, Router } from '@angular/router';
import { fadeInOutAnimation } from '../../animations/animations.module';
import { Hash } from 'crypto';
import { get } from 'http';


@Component({
  selector: 'app-form-edit-redirect',
  standalone: true,
  imports: [IconWhatsappComponent,IconTelegramComponent,SharedModule,ReactiveFormsModule,CommonModule,FormsModule],
  animations:[fadeInOutAnimation],
  templateUrl: './form-edit-redirect.component.html',
  styleUrl: './form-edit-redirect.component.scss'
})
export class FormEditRedirectComponent implements OnInit {
  editedLinks = []
  redirectHash:string = this.activatedRoute.snapshot.params['hash_redirect']

  disableEditNome:boolean = true;
  submitted:boolean = false

  nomeRedirecionador!:string
  formStep:string = 'init'
  getPlataforma!:string
  redirectName!:string
  prioridade:string = 'whatsapp,telegram'

  submitData:any = []
  editData:any = {'whatsappData':[],'telegramData':[]}
  redirectData!:any
  radioMenu!:any

  createLinkForm!:FormGroup

  constructor
  (
    private formBuilder:FormBuilder,
    private api:RedirectifyApiService,
    private router:Router,
    private activatedRoute:ActivatedRoute,

  )
  {


    this.createLinkForm = this.formBuilder.group({
      link:['',[Validators.required]],
      nome:['',[Validators.required]],
      mensagem:[''],
      plataforma:[],
      id:[]
    })
  }



  async ngOnInit() {
    this.redirectData = await this.getRedirectData()
    this.nomeRedirecionador = this.redirectData.redirecionador.nome
    for(let item of this.redirectData.links){
      if(item.plataforma == 'whatsapp' ){
        this.editData['whatsappData'].push(item)
      } else{
        this.editData['telegramData'].push(item)
      }
    }
  }

  async buttonCardEvent(event:string,data:any){
    console.log(data)
    if(event == 'editar'){
      if(data.plataforma == 'whatsapp'){
        this.formStep = 'editW'
        const link = data.link;
        const matches = link.match(/https?:\/\/wa\.me\/\+(\d+)(\?text=(.*))?/);

        let phoneNumber = '';
        let message = '';

        if (matches) {
          phoneNumber = matches[1];
          message = matches[3] || ''; // Se o texto não estiver presente, será uma string vazia
        }
        if(phoneNumber){
          this.createLinkForm.controls['link'].setValue(phoneNumber)
        }
        if(message){
          this.createLinkForm.controls['mensagem'].setValue(decodeURI(message))
        }
        this.createLinkForm.controls['nome'].setValue(data.nome)
        this.createLinkForm.controls['plataforma'].setValue(data.plataforma)
        this.createLinkForm.controls['id'].setValue(data.id)
        console.log(this.createLinkForm.getRawValue())
      } else{
        this.formStep = 'editT'
        this.createLinkForm.controls['link'].setValue(data.link)
        this.createLinkForm.controls['nome'].setValue(data.nome)
        this.createLinkForm.controls['plataforma'].setValue(data.plataforma)
        this.createLinkForm.controls['id'].setValue(data.id)
        console.log(this.createLinkForm.getRawValue())
      }
    }
    else if(event =='deletar'){
      try{
        const res = await this.api.deleteLinkInRedirect(this.redirectHash,data.id)
        if (res.status == 200){
          await this.ngOnInit()
        }
      }
      catch(error){

      }
    }
  }

  async getRedirectData(){
    return await this.api.getRedirect(this.redirectHash)

  }

  editDataEmpty(){
    return Object.keys(this.editData).length?true:false
  }

 async addContact(plataforma:string){
    if (plataforma == 'whatsapp'){
      let formData = [this.createLinkForm.getRawValue()]
      formData[0].plataforma = plataforma
      for (let item of formData) {
        // Obter o valor correspondente à chave atual
        this.submitData = [{
          nome: `${item.nome ? item.nome : '+' + item.link}`,
          link: `https://wa.me/+${item.link}${item.mensagem ? `?text=${encodeURIComponent(item.mensagem)}` : ""}`,
          plataforma: item.plataforma
        }]
      }
      try{
        const resApi = await this.api.addLinkToRedirect(this.redirectHash,this.submitData)
        if(resApi.status == 201){
          this.createLinkForm.reset()
          this.ngOnInit()
          this.formStep = 'init'
        }

      } catch (error) {
        console.log(error)
      }
  }
}

  getContacts(plataforma:string){
    this.getPlataforma = plataforma
    this.formStep = 'getContacts'
  }

 async onSubmit(){
    for(let item of this.editData['whatsappData']){
      this.submitData.push(
        {
        nome:`${item.nome?item.nome:'+'+item.link}`,
        link:`https://wa.me/+${item.link}${item.mensagem ? `?text=${encodeURIComponent(item.mensagem)}` : ""}`,
        plataforma:item.plataforma
        }
      )
    }
    for(let item of this.editData['telegramData']){
      this.submitData.push(
        {
        nome:`${item.nome?item.nome:item.link}`,
        link:item.link,
        plataforma:item.plataforma
        }
      )
    }
    try{
      const resApi = await this.api.addLinkToRedirect(this.redirectHash,this.submitData)
      if (resApi.status == 200) {
        this.router.navigate(['/home/'+this.redirectHash])
      }
      console.log(resApi)
    } catch (error){
      console.log(error)
    }
  }

  async saveEdits(){
    let linkData = this.createLinkForm.getRawValue()
    if(linkData.plataforma == 'whatsapp'){
      for(let item of linkData){
        linkData =
          {
          nome:`${item.nome?item.nome:'+'+item.link}`,
          link:`https://wa.me/+${item.link}${item.mensagem ? `?text=${encodeURIComponent(item.mensagem)}` : ""}`,
          plataforma:item.plataforma
          }
      }
    }
    console.log(linkData)
    // const resSaveEdits = await this.api.addLinkToRedirect(this.redirectHash,linkData)
    // if (resSaveEdits.status == 200){
    //   this.formStep = 'init'
    // }
  }
}
