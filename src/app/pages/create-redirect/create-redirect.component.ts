import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { FormCreateRedirectComponent } from '../../shared/form-create-redirect/form-create-redirect.component';
import { RouterModule } from '@angular/router';


@Component({
  selector: 'app-create-redirect',
  standalone: true,
  imports: [SharedModule,FormCreateRedirectComponent,RouterModule],
  templateUrl: './create-redirect.component.html',
  styleUrl: './create-redirect.component.scss'
})
export class CreateRedirectComponent {

}
