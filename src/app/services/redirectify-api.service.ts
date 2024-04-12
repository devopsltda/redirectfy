import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { lastValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RedirectifyApiService {
  // Rotas
  private prefix:string = 'http://localhost:8080'
  private loginRoute:string = `${this.prefix}/usuarios/login`
  private finishSignUpRoute:string = `${this.prefix}/usuarios/criar_permanente/`


  constructor(private http:HttpClient) { }

   async login(email:string,senha:string) {
    const res = await lastValueFrom(this.http.post<any>(this.loginRoute,{email:email,senha:senha},{observe:'response',withCredentials:true}))
    .catch(
      (error)=>
        {
          throw {status:error.status,statusText:error.statusText, error:error.error}
        }
    )

    return res

  }

  async finishSignUp(hash:string,data_de_nascimento:Date,senha_nova:string){

      const response = await lastValueFrom(this.http.post<any>(`${this.finishSignUpRoute}${hash}`,{data_de_nascimento:data_de_nascimento,senha_nova:senha_nova},{observe:'response'}))
      .catch((error)=>{console.log(error);throw {status:error.status,statusText:error.statusText, error:error.error}})

      return {status:response.status,statusText:response.statusText,ok:response.ok,message:response.body}
  }

}
