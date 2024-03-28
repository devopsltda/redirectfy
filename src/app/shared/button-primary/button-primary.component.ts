import { Component, Input } from '@angular/core';
import { ButtonModule } from 'primeng/button';
@Component({
  selector: 'app-button-primary',
  standalone: true,
  imports: [ButtonModule],
  templateUrl: './button-primary.component.html',
  styleUrl: './button-primary.component.scss'
})
export class ButtonPrimaryComponent {

  @Input() title:string = 'title'
  @Input() type!:string
}
