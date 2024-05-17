import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { NavbarComponent } from '../../shared/navbar/navbar.component';
import { GridComponent } from '../../shared/grid/grid.component';
import { CommonModule } from '@angular/common';
import { SharedModule } from '../../shared/shared.module';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { MessageService } from 'primeng/api';


@Component({
  selector: 'app-my-account',
  standalone: true,
  imports: [SharedModule, RouterModule, NavbarComponent, GridComponent, CommonModule],
  providers:[MessageService],
  templateUrl: './my-account.component.html',
  styleUrl: './my-account.component.scss'
})

export class MyAccountComponent implements OnInit {

  userData: any
  initials: string = '';
  constructor(private api: RedirectifyApiService, private messageService:MessageService,private router:Router, private activatedRoute:ActivatedRoute ) {

  }
  async ngOnInit() {
    await this.getUserData()
  }

  goChangePlan(){
    window.location.href = 'https://app.kirvano.com/'
  }


  async getUserData() {
    this.userData = await this.api.getUser();
    this.userData.cpf = this.userData.cpf ? this.userData.cpf.replace(/\D/g, '').replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, '$1.$2.$3-$4') : this.userData.cpf;
    const data = new Date(this.userData.criado_em);
    const dataNascimento = new Date(this.userData.data_de_nascimento);
    const options: Intl.DateTimeFormatOptions = { day: '2-digit', month: 'long', year: 'numeric' };
    const optionsData: Intl.DateTimeFormatOptions = { day: '2-digit', month: '2-digit', year: 'numeric' };
    this.userData.criado_em = data.toLocaleDateString('pt-BR', options);
    this.userData.data_de_nascimento = dataNascimento.toLocaleDateString('pt-BR', optionsData);
    return this.userData;
  }

  async redifinirSenha(){
    try {
      const res = await this.api.changePasswordUser(this.userData.email)
      if(res.status == 200){
        this.messageService.add({severity:'success', summary:'Redifição de senha solicitada',detail:'Email de redefinição enviado, por favor verifique sua caixa de entrada ou span'})
      }
    } catch(error){
      this.messageService.add({severity:'error', summary:'Falha na Redifição de senha ',detail:'Falha ao redefinir a senha, tente novamente mais tarde, caso persista por favor entre em contato com o suporte'})
    }

  }

}
