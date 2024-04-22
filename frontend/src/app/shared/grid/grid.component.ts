import { CommonModule } from '@angular/common';
import { Component, Input} from '@angular/core';
import { PagesModule } from '../../pages/pages.module';


@Component({
  selector: 'app-grid',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './grid.component.html',
  styleUrl: './grid.component.scss'
})
export class GridComponent {
  @Input() ngClass!:string;
}
