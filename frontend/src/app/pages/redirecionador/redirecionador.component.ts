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
  // Referência dos Deep Links do Telegram: https://core.telegram.org/api/links

  // Prefixo necessário para todos os links do Telegram
  const prefix = "https://t.me/";

  if (link.startsWith(prefix)) {
    const hash = link.slice(prefix.length);

    // Se o link do Telegram for privado, remover o "+" e usar o endpoint para se
    // juntar a grupos privados. CHAT ATRAVÉS DE GRUPOS/CANAIS PRIVADOS
    //
    // telegram doc: https://core.telegram.org/api/links#chat-invite-links
    if (hash.startsWith("+")) {
      return `tg://join?invite=${hash.slice(1)}`;
    }

    // se o link do telegram for um número de telefone(17 dígitos contando espaços e +) BRASILEIRO, 
    // remove o "+" e usa o endpoint para abrir o chat do usuario.
    // CHAT COM USUARIO ATRAVÉS DO NÚMERO DE TELEFONE
    //
    // telegram doc: https://core.telegram.org/api/links#phone-number-links
    if ((hash.length >= 16) && hash.startsWith("+55")) {
      return `tg://resolve?phone=${hash.slice(1)}`;
    }


    // Se não o link do Telegram é considerado público, usar o endpoint para se
    // juntar a grupos públicos sem alterar o hash. 
    // CHAT ATRAVÉS DE NOME DE USUARIOS
    //
    // telegram doc: https://core.telegram.org/api/links#public-username-links
    return `tg://resolve?domain=${hash}`;
  }

  // Caso o link não inicie com o prefixo necessário, ele deve ser devolvido da
  // maneira como foi enviado.
  return link;
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
            if(this.linkTelegram){
              window.location.href = this.telegramLinkToHook(this.linkTelegram)
            }
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

