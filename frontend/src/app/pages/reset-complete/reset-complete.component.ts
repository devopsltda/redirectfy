import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { RouterModule } from '@angular/router';


@Component({
  selector: 'app-reset-complete',
  standalone: true,
  imports: [CommonModule,SharedModule,RouterModule],
  templateUrl: './reset-complete.component.html',
  styleUrl: './reset-complete.component.scss'
})
export class ResetCompleteComponent {

}
