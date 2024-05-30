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
  dataEdit:any
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

  async buttonCardEvent(event:string,data:any){
    if(event == 'editar'){
      if(data.plataforma == 'whatsapp'){
        this.whatsappForm.controls['link'].setValue(data.link);
        this.whatsappForm.controls['mensagem'].setValue(data.mensagem);
        this.whatsappForm.controls['nome'].setValue(data.nome)
        this.goEdit(data)
        this.formStep = 'editW'
      }
      else if(data.plataforma == 'telegram'){
       this.telegramForm.controls['link'].setValue(data.link);
       this.telegramForm.controls['nome'].setValue(data.nome)
       this.goEdit(data)
       this.formStep = 'editT'
     }
    } else if (event == 'deletar') {
      if (data.plataforma == 'whatsapp') {
          this.createData['whatsappData'] = this.createData['whatsappData'].filter((item: any) => item.nome !== data.nome);
          if(this.createData['whatsappData'].length == 0 ){
            delete this.createData['whatsappData']
            this.prioridade = 'telegram,whatsapp'
          }
      } else {
        this.createData['telegramData'] = this.createData['telegramData'].filter((item: any) => item.nome !== data.nome);
        if(this.createData['telegramData'].length == 0 ){
          delete this.createData['telegramData']
          this.prioridade = 'whatsapp,telegram'
        }
      }
  }


  }

  goEdit(data:any){
      this.dataEdit = data
  }

  saveEdit(plataforma: string) {
    if (plataforma == 'whatsapp') {
        const index = this.createData['whatsappData'].findIndex((item: any) => item.nome === this.dataEdit.nome);
        if (index !== -1) {
            this.createData['whatsappData'][index].link = this.whatsappForm.get('link')?.value;
            this.createData['whatsappData'][index].mensagem = this.whatsappForm.get('mensagem')?.value;
            this.createData['whatsappData'][index].nome = this.whatsappForm.get('nome')?.value;
        }
    } else if (plataforma == 'telegram') {
        const index = this.createData['telegramData'].findIndex((item: any) => item.nome === this.dataEdit.nome);
        if (index !== -1) {
            this.createData['telegramData'][index].link = this.telegramForm.get('link')?.value;
            this.createData['telegramData'][index].nome = this.telegramForm.get('nome')?.value;
        }
    }
    this.formStep = 'init';
    this.dataEdit = null;
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

    if (this.redirectName.length < 3 || this.redirectName.length > 20) {
      this.messageService.add({summary: "Falha ao Criar Redirecionador", detail: 'Por favor, insira um nome entre 3 e 20 caracteres', severity: 'error'});
    } else {
      try {
        const resApi = await this.api.createRedirect(this.redirectName, this.prioridade, this.submitData);
        if (resApi.status === 201) {
          this.router.navigate(['/home']);
        }
      }
      catch (error) {
        if (typeof error === 'object' && error !== null) {
          if((error as any).status === 402){
            this.messageService.add({summary: "Limite de criação excedido!", detail: 'Oops! Parece que você atingiu o limite de redirecionadores. Faça um upgrade do seu plano de assinatura ou remova redirecionadores já existentes', severity: 'info'});
          }
        }else{
          this.messageService.add({summary: "Falha ao Criar Redirecionador", detail: 'Ocorreu um erro ao criar o redirecionador, ação não executada', severity: 'error'});
        }
      }
    }
  }
}
