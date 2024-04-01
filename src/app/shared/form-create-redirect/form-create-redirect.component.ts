import { Component } from '@angular/core';
import { IconWhatsappComponent } from '../icon-whatsapp/icon-whatsapp.component';
import { IconTelegramComponent } from '../icon-telegram/icon-telegram.component';
import { SharedModule } from '../shared.module';
import { AnimationsModule, fadeInOutAnimation } from '../../animations/animations.module';

@Component({
  selector: 'app-form-create-redirect',
  standalone: true,
  imports: [IconWhatsappComponent,IconTelegramComponent,SharedModule,AnimationsModule],
  animations:[fadeInOutAnimation],
  templateUrl: './form-create-redirect.component.html',
  styleUrl: './form-create-redirect.component.scss'
})
export class FormCreateRedirectComponent {
  formStep:string = 'init'
}
