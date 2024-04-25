import { Component, OnInit } from '@angular/core';
import { RouterModule } from '@angular/router';
import { NavbarComponent } from '../../shared/navbar/navbar.component';
import { GridComponent } from '../../shared/grid/grid.component';
import { CommonModule } from '@angular/common';
import { SharedModule } from '../../shared/shared.module';
import { RedirectifyApiService } from '../../services/redirectify-api.service';


@Component({
  selector: 'app-my-account',
  standalone: true,
  imports: [SharedModule, RouterModule, NavbarComponent, GridComponent, CommonModule],
  templateUrl: './my-account.component.html',
  styleUrl: './my-account.component.scss'
})

export class MyAccountComponent implements OnInit {

  userData: any
  initials: string = '';
  constructor(private api: RedirectifyApiService) {

  }
  async ngOnInit() {
    await this.getUserData()
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

}
