import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { GridComponent } from '../../shared/grid/grid.component';
import { SharedModule } from '../../shared/shared.module';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-recovery-password',
  standalone: true,
  imports: [CommonModule,SharedModule,RouterModule],
  templateUrl: './recovery-password.component.html',
  styleUrl: './recovery-password.component.scss'
})
export class RecoveryPasswordComponent {

}
