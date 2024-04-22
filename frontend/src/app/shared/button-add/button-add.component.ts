import { Component, Input } from '@angular/core';
import { IconAddCircledComponent } from '../icon-add-circled/icon-add-circled.component';

@Component({
  selector: 'app-button-add',
  standalone: true,
  imports: [IconAddCircledComponent],
  templateUrl: './button-add.component.html',
  styleUrl: './button-add.component.scss'
})
export class ButtonAddComponent {
  @Input() title:string = 'title'
  @Input() type!:string
}
