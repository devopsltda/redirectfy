import { Component, OnInit } from '@angular/core';
import { IconWhatsappComponent } from '../icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../icon-telegram/icon-telegram.component';
import { SharedModule } from '../shared.module';

import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { ActivatedRoute, Router } from '@angular/router';
import { fadeInOutAnimation } from '../../animations/animations.module';


@Component({
  selector: 'app-form-edit-redirect',
  standalone: true,
  imports: [IconWhatsappComponent,IconTelegramComponent,SharedModule,ReactiveFormsModule,CommonModule,FormsModule],
  animations:[fadeInOutAnimation],
  templateUrl: './form-edit-redirect.component.html',
  styleUrl: './form-edit-redirect.component.scss'
})
export class FormEditRedirectComponent implements OnInit {
  disableEditNome:boolean = true;
  submitted:boolean = false
  nomeRedirecionador!:string
  formStep:string = 'init'
  getPlataforma!:string
  redirectName!:string
  prioridade:string = 'whatsapp,telegram'
  submitData:any = []
  editData:any = {'whatsappData':[],'telegramData':[]}
  createLinkForm!:FormGroup
  redirectData!:any
  redirectHash:string = this.activatedRoute.snapshot.params['hash_redirect']

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
      plataforma:[]
    })
  }

  teste(){
    console.log('oi')
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

  async getRedirectData(){
    return await this.api.getRedirect(this.redirectHash)

  }

  editDataEmpty(){
    return Object.keys(this.editData).length?true:false
  }

  addContact(plataforma:string){
    this.submitted = true
    if(this.createLinkForm.valid){
      this.submitted = false
      if(plataforma == 'whatsapp'){
        const data = this.createLinkForm.getRawValue()
        data['plataforma'] = plataforma
        if(this.editData['whatsappData'] == undefined){
          this.editData['whatsappData'] = [data]
          this.createLinkForm.reset()
          this.formStep = 'init'
        } else{
          this.editData['whatsappData'].push(data)
          this.createLinkForm.reset()
          this.formStep = 'init'
        }

        console.log(this.editData['whatsappData'])
      } else if(plataforma == 'telegram'){

        const data = this.createLinkForm.getRawValue()
        data['plataforma'] = plataforma

        if(this.editData['telegramData'] == undefined){

          this.editData['telegramData'] = [data]
          this.createLinkForm.reset()
          this.formStep = 'init'
        } else{

          this.editData['telegramData'].push(data)
          this.createLinkForm.reset()
          this.formStep = 'init'
        }
        console.log(this.editData)
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
