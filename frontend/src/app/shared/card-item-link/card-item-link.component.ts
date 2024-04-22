import { Component, EventEmitter, Input, Output } from '@angular/core';
import { ButtonCopyComponent } from '../button-copy/button-copy.component';
import { ButtonShareComponent } from '../button-share/button-share.component';
import { ButtonDeleteComponent } from '../button-delete/button-delete.component';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-card-item-link',
  standalone: true,
  imports: [ButtonCopyComponent,ButtonShareComponent,ButtonDeleteComponent,RouterLink],
  templateUrl: './card-item-link.component.html',
  styleUrl: './card-item-link.component.scss'
})
export class CardItemLinkComponent {
  @Input() cardTitle:string = 'Card Title'
  @Input() hash:string = ''
  @Output() botãoClicado = new EventEmitter()

  onClick(event:string){
    this.botãoClicado.emit(event)
  }
}
