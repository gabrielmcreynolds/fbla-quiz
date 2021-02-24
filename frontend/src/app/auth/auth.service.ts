import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, of, Subject } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthData } from './auth-data/auth-data.module';
import { User } from './user';
import { AuthStatus } from './auth-status.enum';
import { catchError, mapTo, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(private http: HttpClient, private router: Router) {
    this.authStatusListener.subscribe(
      (value) => (this.isAuthenticated = value)
    );
    this.user$ = new BehaviorSubject<User>(null);
    setTimeout(() => this.initUser(), 100);
  }

  private accessToken: string;
  private authStatusListener = new Subject<boolean>();
  private isAuthenticated = false;
  private user$: Subject<User>;

  private static saveAuthData(refreshToken: string, accessToken: string): void {
    localStorage.setItem('refreshToken', refreshToken);
    localStorage.setItem('accessToken', accessToken);
  }

  private static getRefreshToken(): string {
    return localStorage.getItem('refreshToken');
  }

  setUser(user: User): void {
    this.user$.next(user);
  }

  initUser(): void {
    if (localStorage.getItem('refreshToken') != null) {
      this.http.get<{ user: User }>('users').subscribe(
        (data) => {
          if (data) {
            this.user$.next(data.user);
            this.router.navigate(['/dashboard']);
          }
        },
        () => {
          this.logout();
        }
      );
    }
    this.user$.next(null);
    this.router.navigate(['/']);
  }

  private clearAuthData(): void {
    localStorage.clear();
    this.user$.next(null);
    this.accessToken = null;
  }

  createUser(authData: AuthData): Observable<AuthStatus> {
    return this.http
      .post<{
        accessToken: string;
        refreshToken: string;
        user: User;
      }>('users/signup', authData)
      .pipe(
        tap((response) => {
          if (response != null) {
            this.setAccessToken(response.accessToken);
            AuthService.saveAuthData(
              response.refreshToken,
              response.accessToken
            );
            this.user$.next(response.user);
          }
        }),
        mapTo(AuthStatus.LoggedIn),
        catchError((err) => {
          if (err.status === 409) {
            return of(AuthStatus.IncorrectEmail);
          } else {
            return of(AuthStatus.UnknownError);
          }
        })
      );
  }

  login(email: string, password: string): Observable<AuthStatus> {
    const authData: AuthData = { email, password };
    return this.http
      .post<{
        accessToken: string;
        refreshToken: string;
        user: User;
      }>('users/login', authData)
      .pipe(
        tap((response) => {
          if (response != null) {
            this.accessToken = response.accessToken;
            AuthService.saveAuthData(
              response.refreshToken,
              response.accessToken
            );
            this.user$.next(response.user);
            this.authStatusListener.next(true);
            this.router.navigate(['/dashboard']);
          }
        }),

        mapTo(AuthStatus.LoggedIn),

        catchError((err) => {
          console.log(`Found Error: ${err}`);
          if (err.status === 404) {
            return of(AuthStatus.IncorrectEmail);
          } else if (err.status === 401) {
            return of(AuthStatus.IncorrectPassword);
          } else {
            return of(AuthStatus.UnknownError);
          }
        })
      );
  }

  refresh(): Observable<{ accessToken: string; message: string }> {
    return this.http.post<{ accessToken: string; message: string }>(
      `users/refresh`,
      { refreshToken: AuthService.getRefreshToken() }
    );
  }

  getUser = () => this.user$.asObservable();

  getAccessToken(): string {
    if (this.accessToken == null) {
      return localStorage.getItem('accessToken');
    }
    return this.accessToken;
  }

  setAccessToken(token: string): void {
    this.accessToken = token;
    localStorage.setItem('accessToken', token);
  }

  logout(): void {
    const refreshToken = AuthService.getRefreshToken();
    this.http
      .request('delete', 'users/logout', {
        body: {
          refreshToken,
        },
      })
      .subscribe(() => {
        this.clearAuthData();
      });
    this.router.navigate(['/']);
  }
}
