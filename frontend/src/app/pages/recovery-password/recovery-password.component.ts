import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { GridComponent } from '../../shared/grid/grid.component';
import { SharedModule } from '../../shared/shared.module';
import { Router, RouterModule } from '@angular/router';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { FormsModule } from '@angular/forms';
import { MessageService } from 'primeng/api';

@Component({
  selector: 'app-recovery-password',
  standalone: true,
  imports: [CommonModule,SharedModule,RouterModule,FormsModule],
  providers:[MessageService],
  templateUrl: './recovery-password.component.html',
  styleUrl: './recovery-password.component.scss'
})
export class RecoveryPasswordComponent {

  recoveryEmail!:string

  constructor(private api:RedirectifyApiService,private router:Router, private messageService:MessageService){

  }

  async recoveryPassword(){
    try {
      const res = await this.api.changePasswordUser(this.recoveryEmail)
      console.log(res)
      if (res.status == 200){
        this.router.navigate(['/recoverySend'])
      }
    } catch (error) {
      return this.messageService.add({severity:'erro',summary:'Falha na ação',detail:'Fala ao solicitar redefinição de senha'})
    }
  }
}
