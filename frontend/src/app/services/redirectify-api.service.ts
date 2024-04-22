import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { lastValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class RedirectifyApiService {
  // Rotas
  private prefix: string = '/api';
  private loginRoute: string = `${this.prefix}/u/login`;
  private logoutRoute: string = `${this.prefix}/u/logout`;
  private finishSignUpRoute: string = `${this.prefix}/kirvano/to_user/`;
  private getAllRedirectsRoute: string = `${this.prefix}/r`;
  private createRedirectRoute: string = `${this.prefix}/r`;
  private deleteRedirectRoute: string = `${this.prefix}/r`;
  private getRedirectRoute: string = `${this.prefix}/r`;
  private addLinkToRedirectRoute: string = `${this.prefix}/r`;
  private deleteLinkInRedirectRoute: string = `${this.prefix}/r`;
  private getUserRoute: string = `${this.prefix}/u`
  private updateLinkInRedirectRoute:string = `${this.prefix}/r`
  constructor(private http: HttpClient, private cookies: CookieService) {}

  // Usuário
  async getUser(){
    const resGetUser = await lastValueFrom(this.http.get(this.getUserRoute,{withCredentials:true,observe:'response'}))
    .catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resGetUser.body
  }

  // redirecionadores
  async disableLinkInRedirect(hash:string,idLink:number){
    const resGetRedirect = await lastValueFrom(this.http.patch(`${this.updateLinkInRedirectRoute}/${hash}/links/${idLink}/disable`,{},{withCredentials:true,observe:'response'}))
    .catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resGetRedirect
  }

  async enableLinkInRedirect(hash:string,idLink:number){
    const resGetRedirect = await lastValueFrom(this.http.patch(`${this.updateLinkInRedirectRoute}/${hash}/links/${idLink}/enable`,{},{withCredentials:true,observe:'response'}))
    .catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resGetRedirect
  }

  async updateLinkInRedirect(hash:string,idLink:number,data:any){
    const resGetRedirect = await lastValueFrom(this.http.patch(`${this.updateLinkInRedirectRoute}/${hash}/links/${idLink}`,data,{withCredentials:true,observe:'response'}))
    .catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resGetRedirect
  }
  async deleteLinkInRedirect(hash:string,idLink:number){
    const resGetRedirect = await lastValueFrom(this.http.delete(`${this.deleteLinkInRedirectRoute}/${hash}/links/${idLink}`,{withCredentials:true,observe:'response'}))
    .catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resGetRedirect
  }


  async addLinkToRedirect(hash:string,link:any){
    const resGetRedirect = await lastValueFrom(this.http.post(`${this.addLinkToRedirectRoute}/${hash}/links`,{links:link},{withCredentials:true,observe:'response'}))
    .catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resGetRedirect
  }

  async getRedirect(hash:string){
    const resGetRedirect = await lastValueFrom(this.http.get(`${this.getRedirectRoute}/${hash}`,{withCredentials:true}))
    .catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resGetRedirect
  }

  async deleteRedirect(codigoHash:string){
    const resDeleteRedirect = await lastValueFrom(this.http.delete(`${this.deleteRedirectRoute}/${codigoHash}`,{withCredentials:true}))
    .catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resDeleteRedirect
  }
  async createRedirect(nome: string, ordem_de_redirecionamento: string,links:{link:string,nome:string,plataforma:string}[]) {

    const resCreateRedirect = await lastValueFrom(
      this.http.post<any>(
        this.createRedirectRoute,
        {
          nome: nome,
          links:links,
          ordem_de_redirecionamento: ordem_de_redirecionamento,
        },
        { withCredentials: true,observe:'response' }
      )
    ).catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return resCreateRedirect

  }

  async getAllRedirects() {
    const res = await lastValueFrom(
      this.http.get<any>(this.getAllRedirectsRoute, { withCredentials: true })
    ).catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });
    return res;
  }

  async login(email: string, senha: string) {
    const res = await lastValueFrom(
      this.http.post<any>(
        this.loginRoute,
        { email: email, senha: senha },
        { observe: 'response', withCredentials: true }
      )
    ).catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });
    return res;
  }
  //  user
  async logout(){
    const res = await lastValueFrom(
      this.http.post<any>(
        this.logoutRoute,'',
        { observe: 'response', withCredentials: true }
      )
    ).catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });
    return res;
  }

  async finishSignUp(
    hash: string,
    data_de_nascimento: Date,
    senha_nova: string
  ) {
    const response = await lastValueFrom(
      this.http.post<any>(
        `${this.finishSignUpRoute}${hash}`,
        { data_de_nascimento: data_de_nascimento, senha: senha_nova },
        { observe: 'response' }
      )
    ).catch((error) => {
      throw {
        status: error.status,
        statusText: error.statusText,
        error: error.error,
      };
    });

    return {
      status: response.status,
      statusText: response.statusText,
      ok: response.ok,
      message: response.body,
    };
  }
}
