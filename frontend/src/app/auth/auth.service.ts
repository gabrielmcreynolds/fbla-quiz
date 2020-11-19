import {Injectable} from '@angular/core';
import {Observable, Subject} from 'rxjs';
import {HttpClient} from '@angular/common/http';
import {Router} from '@angular/router';
import {AuthData} from './auth-data/auth-data.module';
import {User} from './user';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient, private router: Router) {
    this.authStatusListener.subscribe((value) => this.isAuthenticated = value);
    this.user$ = new Subject<User>();
  }

  private accessToken: string;
  private authStatusListener = new Subject<boolean>();
  private isAuthenticated = false;
  private user$: Subject<User>;

  private static saveAuthData(
    refreshToken: string,
    accessToken: string,
  ): void {
    localStorage.setItem('refreshToken', refreshToken);
    localStorage.setItem('accessToken', accessToken);
  }

  private static getRefreshToken(): string {
    return localStorage.getItem('refreshToken');
  }

  private clearAuthData(): void {
    localStorage.clear();
    this.user$.next(null);
    this.accessToken = null;
  }

  login(email: string, password: string): void {
    const authData: AuthData = {email, password};
    this.http.post<{
      accessToken: string; refreshToken: string;
      user: User
    }>('users/login', authData).subscribe(
      (response) => {
        if (response) {
          this.accessToken = response.accessToken;
          AuthService.saveAuthData(response.refreshToken, response.accessToken);
          this.user$.next(response.user);
          this.authStatusListener.next(true);
          this.router.navigate(['/dashboard']);
        }
      }, (error => {
        this.authStatusListener.next(false);
      })
    );
  }

  refresh(): Observable<{ accessToken: string, message: string }> {
    return this.http.post<{ accessToken: string, message: string }>(
      `users/refresh`, {refreshToken: AuthService.getRefreshToken()});
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
    this.http.request('delete', 'users/logout', {
      body: {
        refreshToken,
      }
    }).subscribe(data => {
      this.clearAuthData();
    });
    this.router.navigate(['/']);
  }
}
