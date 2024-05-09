import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-devops-brand',
  standalone: true,
  imports: [],
  templateUrl: './devops-brand.component.html',
  styleUrl: './devops-brand.component.scss'
})
export class DevopsBrandComponent {

  @Input() class:string = 'w-5 h-5 fill-black'

}
