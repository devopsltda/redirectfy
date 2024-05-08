import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormsModule } from '@angular/forms';


@Component({
  selector: 'app-card-contact-mobile',
  standalone: true,
  imports: [FormsModule, CommonModule],
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

  ngOnInit(): void {
    document.body.addEventListener('click', () => {
      if (this.menuIsOpen) {
        this.menuIsOpen = false;
      }
    });
  }

  toggleMenu(event: Event) {
    this.menuIsOpen = !this.menuIsOpen;
    event.stopPropagation(); 
  }
  
  handleButtonClick(event: Event) {
    event.stopPropagation();
  }
}