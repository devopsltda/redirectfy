import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { ConfirmationService } from 'primeng/api';
import { SharedModule } from '../../shared/shared.module';
import { RedirecionadorAnimation} from '../../animations/animations.module';
import { IconWhatsappComponent } from '../../shared/icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../../shared/icon-telegram/icon-telegram.component';

@Component({
  selector: 'app-redirecionador',
  standalone: true,
  imports: [ IconWhatsappComponent,IconTelegramComponent,SharedModule,ConfirmDialogModule],
  animations:[RedirecionadorAnimation],
  providers:[ConfirmationService],
  templateUrl: './redirecionador.component.html',
  styleUrl: './redirecionador.component.scss'
})
export class RedirecionadorComponent implements OnInit {

  redirectHash:string = this.activatedRoute.snapshot.params['hash']
  data:any
  link1!:string
  trustedLink!:SafeUrl
  link2!:string
  isLoading:boolean = true
  linkTelegram = "tg://join?invite=mjJCtz57OtU5MzBh";
  linkWhatsapp = "whatsapp://send/?phone=5598991085854";
  constructor
  (
    private activatedRoute:ActivatedRoute,
    private api:RedirectifyApiService,
    private sanitazer:DomSanitizer,
    private confirmationService:ConfirmationService,
  ){

  }

  async ngOnInit() {
    this.data = await this.api.getLinksRedirect(this.redirectHash)
    setTimeout(()=>{
      this.openDialog()
    },3000)

  }

  goLinkTelegram(){
    return  window.location.href = this.linkTelegram;
  }

  goLinkWhatsapp(){
    return  window.location.href = this.linkWhatsapp;
  }

  openDialog(){
    const linkTelegram = "tg://join?invite=+a1PN4mwBrTFkY2Ex";
    const linkWhatsapp = "whatsapp://send/?phone=5598991085854";
    this.confirmationService.confirm({
      message: 'Abrir whatsapp e iniciar a conversa?',
      accept: () => {
        window.location.href = linkWhatsapp;
        this.isLoading = false
      },
      reject: () => {
        window.location.href = linkTelegram;
        this.isLoading = false
      }
    })
  }
}
