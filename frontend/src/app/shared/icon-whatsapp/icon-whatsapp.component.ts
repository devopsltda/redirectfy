import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { CommonEngine } from '@angular/ssr';

@Component({
  selector: 'app-icon-whatsapp',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './icon-whatsapp.component.html',
  styleUrl: './icon-whatsapp.component.scss'
})
export class IconWhatsappComponent {

  @Input() ngClass!:string

}
