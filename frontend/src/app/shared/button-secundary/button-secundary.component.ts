import { Component, Input } from '@angular/core';
@Component({
  selector: 'app-button-secundary',
  standalone: true,
  imports: [],
  templateUrl: './button-secundary.component.html',
  styleUrl: './button-secundary.component.scss'
})
export class ButtonSecundaryComponent {
  @Input() title:string = 'title'
  @Input() type!:string
}
