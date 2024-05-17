import { Component, ComponentFactoryResolver, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { ConfirmDialogModule } from 'primeng/confirmdialog';

import { ConfirmationService } from 'primeng/api';
import { SharedModule } from '../../shared/shared.module';
import { RedirecionadorAnimation} from '../../animations/animations.module';
import { IconWhatsappComponent } from '../../shared/icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../../shared/icon-telegram/icon-telegram.component';
import { CommonModule } from '@angular/common';


@Component({
  selector: 'app-redirecionador',
  standalone: true,
  imports: [ IconWhatsappComponent,IconTelegramComponent,SharedModule,ConfirmDialogModule,CommonModule],
  animations:[RedirecionadorAnimation],
  providers:[ConfirmationService],
  templateUrl: './redirecionador.component.html',
  styleUrl: './redirecionador.component.scss'
})
export class RedirecionadorComponent implements OnInit {

  redirectHash:string = this.activatedRoute.snapshot.params['hash']
  data:any

  isLoading:boolean = true
  linkTelegram!:string
  linkWhatsapp!:string
  IsAccepted:boolean = true
  isVisible:boolean = true

  constructor
  (
    private activatedRoute:ActivatedRoute,
    private api:RedirectifyApiService,
    private confirmationService:ConfirmationService,
  ){
  }

  async ngOnInit() {
    this.data = await this.api.getToLinksRedirect(this.redirectHash)
    console.log(this.data)
    if(this.data.body.links?.[0]?.plataforma == 'whatsapp'){
      this.linkWhatsapp = this.data.body.links?.[0].link
      this.linkTelegram = this.data.body.links?.[1].link
    } else {
      this.linkWhatsapp = this.data.body.links?.[1].link
      this.linkTelegram = this.data.body.links?.[0].link
    }
    this.openDialog()

  }

  goLinkTelegram(){
    return  window.location.href = this.linkTelegram;
  }

  goLinkWhatsapp(){
    console.log(this.linkWhatsapp)
    return  window.location.href = this.linkWhatsapp;
  }

  whatsappLinkToHook(link: string): string {
    // Extrair o número de telefone do link
    const phoneRegex = /\+(\d+)/;
    const phoneMatch = link.match(phoneRegex);
    const phone = phoneMatch ? phoneMatch[1] : '';

    // Extrair o texto do link
    const textRegex = /text=([^&]*)/;
    const textMatch = link.match(textRegex);
    const newText = textMatch ? decodeURIComponent(textMatch[1]) : '';

    // Criar o novo link do WhatsApp com o número de telefone e o texto
    const whatsappLink = `whatsapp://send/app/?phone=${phone}&text=${encodeURI(newText)}`;

    return whatsappLink;
}

telegramLinkToHook(link: string): string {
  // Expressão regular para encontrar tudo após a última barra do "https://t.me/"
  const regex = /https:\/\/t\.me\/\+?([^/]+)$/;
  // Encontra o que está após a última barra do "https://t.me/" no link
  const match = link.match(regex);
  if (match) {
      // Remove o sinal de mais (+) se estiver presente no código de convite
      const inviteCode = match[1].startsWith('+') ? match[1].slice(1) : match[1];
      // Retorna o link do Telegram com o código de convite modificado
      return `tg://join?invite=${inviteCode}`;
  } else {
      // Se não encontrar "https://t.me/" no link, retorna o link original
      return link;
  }
}

   openDialog(){
      if(this.data?.body.redirecionador.ordem_de_redirecionamento == 'whatsapp,telegram'){
        this.confirmationService.confirm({
          header:'Redirecionando para Whatsapp',
          message: `Abrir whatsapp e iniciar a conversa com ${this.data.body?.redirecionador.nome} ?`,

          accept: () => {
            window.location.href = this.whatsappLinkToHook(this.linkWhatsapp)
            this.isLoading = false
          },
          reject: () => {
            this.isLoading = false
            window.location.href = this.telegramLinkToHook(this.linkTelegram)
          }
        })
      } else {
        this.confirmationService.confirm({
          header:'Redirecionando para Telegram',
          message: `Abrir telegram e iniciar a conversa com ${this.data.body?.redirecionador.nome}?`,
          accept: () => {
            this.isLoading = false
            window.location.href = this.telegramLinkToHook(this.linkTelegram)
          },
          reject: () => {
            this.isLoading = false
            window.location.href = this.whatsappLinkToHook(this.linkWhatsapp)
          }
        })
      }
    }
  }

