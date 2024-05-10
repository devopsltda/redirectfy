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
    document.body.addEventListener('click', (event) => {
      if (this.menuIsOpen && event.target) {
        const target = event.target as HTMLElement;
        if (!target.closest('.menu') && !target.closest('.menu-button')) {
          this.menuIsOpen = false;
        }
      }
    });
  }
  
  toggleMenu(event: Event, cardName: string) {
    const menuButtonId = `cardMenuButton${cardName}`;
    const otherMenus = document.querySelectorAll('.menu');
    otherMenus.forEach((menu: Element) => {
      const menuElement = menu as HTMLElement;
      if (menuElement.id !== menuButtonId) {
        menuElement.classList.remove('flex');
        menuElement.classList.add('hidden');
      }
    });

    this.menuIsOpen = !this.menuIsOpen;
    event.stopPropagation(); 
  }

  handleButtonClick(event: Event) {
    event.stopPropagation();
  }
}
