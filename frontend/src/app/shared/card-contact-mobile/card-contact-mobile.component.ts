import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormsModule } from '@angular/forms';


@Component({
  selector: 'app-card-contact-mobile',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './card-contact-mobile.component.html',
  styleUrl: './card-contact-mobile.component.scss'
})
export class CardContactMobileComponent {
  @Input() enableButton:boolean = true;
  @Output() cardEvent = new EventEmitter()
  @Input() cardName = 'Card Name'
  @Input() isTelegram:boolean = false
  @Input() ativo:boolean = true
  menuIsOpen:boolean = false
  menuRadio!:boolean

  cardEventEmmiter(event:string){
    this.cardEvent.emit(event)
  }

}
