import {Injectable} from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor, HttpErrorResponse,
} from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
import {AuthService} from './auth.service';
import {catchError, mergeMap} from 'rxjs/operators';
import {environment} from 'src/environments/environment';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  constructor(private authService: AuthService) {
  }

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    const authToken = this.authService.getAccessToken();
    let authRequest = request.clone({
      headers: request.headers.set('Authorization', `Bearer ${authToken}`),
      url: environment.apiUrl + request.url,
    });
    return next.handle(authRequest).pipe(
      catchError((error: HttpErrorResponse) => {
        if (error.status === 401 && error.error.message === 'invalid or expired jwt') {
          return this.authService.refresh().pipe(
            mergeMap(data => {
                if (data) {
                  this.authService.setAccessToken(data.accessToken);
                  authRequest = request.clone({
                    headers: request.headers.set('Authorization', `Bearer ${data.accessToken}`),
                    url: environment.apiUrl + request.url,
                  });
                  return next.handle(authRequest);
                } else {
                  this.authService.logout();
                  return throwError(error);
                }
              }
            ));
        }
        return throwError(error);
      })
    );
  }
}
