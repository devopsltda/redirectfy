import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { lastValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class RedirectifyApiService {
  // Rotas
  private prefix: string = 'http://localhost:8080';
  private loginRoute: string = `${this.prefix}/u/login`;
  private finishSignUpRoute: string = `${this.prefix}/kirvano/to_user/`;
  private getAllRedirectsRoute: string = `${this.prefix}/r`;
  private createRedirectRoute: string = `${this.prefix}/r`;
  private createRedirectLinksRoute: string = `${this.prefix}/r`;

  constructor(private http: HttpClient, private cookies: CookieService) {}



  async createRedirect(nome: string, ordem_de_redirecionamento: string,links:{link:string,nome:string,plataforma:string}[]) {

    const resCreateRedirect = await lastValueFrom(
      this.http.post<any>(
        this.createRedirectRoute,
        {
          nome: nome,
          ordem_de_redirecionamento: ordem_de_redirecionamento,
        },
        { withCredentials: true }
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
