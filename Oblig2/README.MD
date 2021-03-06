# IS-105, Obligatorisk oppgave 2

Adrian Nesvik, Edvin Myhren, Endre Grindhaug og Tord Vetle Gjertsen


### Oppgave 1

----------------------------------------------------------------------------------------------------------

#### a) Skriv et Golang program med navn fileinfo.go​.
Programmet ligger i 'IS-105_oppgaver/Oblig2/src/oppgave1', under navnet fileinfo.go.
![in cmd](https://raw.githubusercontent.com/StavenX/IS-105_oppgaver/master/Oblig2/images/FileInfo.png "Shows file-information.")

Programmet returnerer informasjon om en fil. 'info' funksjonen bruker et if statement for å printe ut feilmeldinger dersom programmet ikke har fått et argument fra kommando-linjen. I neste metode blir argumentet assigned til variabelen 'filename'. En feilmelding blir printet dersom det skjer en feil når filen blir åpnet. I neste metode blir informasjon om filens størrelse printet ut i forskjellige måleenheter. Videre sjekkes det om det er 'directory', 'regular', 'append', 'file' og 'symbolic' eller ikke. Svaret blir sjekket gjennom metoder og printet ut gjennom et if/else statement. Til slutt blir 'info' funksjonen kjørt gjennom 'main' funksjonen.

#### b) Lag en kjørbart program (build) av fileinfo.go som kan utføres fra kommandolinje. Programmet skal ta et argument, som skal være et filnavn.
Den kompilerte filen (fileinfo.exe) kan bli funnet i (https://raw.githubusercontent.com/StavenX/IS-105_oppgaver/master/Oblig2/bin/fileinfo.exe).

### Oppgave 2

----------------------------------------------------------------------------------------------------------

Programmet ligger i (Oblig2/src/oppgave2), og heter filecount.go. Den kompilerte filen (filecount.exe) kan bli funnet i (https://raw.githubusercontent.com/StavenX/IS-105_oppgaver/master/Oblig2/bin/filecount.exe).
![in cmd](https://raw.githubusercontent.com/StavenX/IS-105_oppgaver/master/Oblig2/images/FileCount.png "Shows most used runes.")

#### a) Skriv programmer med navn filecount.go som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem “runes” (bokstaver/tegn), som forekommer mest i filen samt hvor mange ganger de telles i den medfølgende filen text.txt. Som forrige oppgave skal filen bygges til et kjørbart program og kunne utføres i terminal med kommandoen: filecount  <filnavn>

Vi lager et map med 'make' funksjonen. Den første metoden i main funksjonen printer ut feilmelding dersom  programmet ikke har fått argument i kommandolinjen. I neste metode åpner filen og teller antall linjer gjennom en scanner og en for loop. Deretter printes resultatet, filen konverteres til string og hver rune blir splittet. I neste funksjon finner vi tegn gjennom en for loop og teller tegnene ved bruk av et if/else statement. I funksjon 'MapSorter' brukes for loops til å sammenligne og finne det mest brukte tegnet og det blir deretter deklarert i et if statement. Tegnet " " blir presentert som "(space)" og det mest brukte tegnet blir printet, slettet fra 'RuneMap' og loopen blir gjentatt til loopen stopper.

### Oppgave 3

----------------------------------------------------------------------------------------------------------

Programmene ligger i 'IS-105_oppgaver/Oblig2/src/oppgave3/', under navnene addtofile.go, addup.go og sumfromfile.go.

#### a) Skriv et program med navn addup.go bestående av to go funksjoner og kommunisere med hverandre gjennom channels
#### Funksjon A: 
'funcA' tar inn to tall. Aksepterer kun tall, og gir feil hvis man skriver bokstaver (eller generelt ikke-aksepterte karakterer)
#### Funksjon B: 
'funcB' summerer to tall skrevet inn i 'funcA', og summerer dem. Deretter legger summen 'sum' inn i channel 'ch2'.
 

#### b) Skriv to programmer (addtofile.go og sumfromfile.go) som kommuniserer med hverandre gjennom en fil.

addtofile.go: Leser inn to tall og skriver disse til en fil.
sumfromfile.go: Leser tall fra fil laget av program A og adderer de to tallene. Summen skrives tilbake til samme fil. Program A leser resultatet fra fil og skriver ut summen til stdout.

#### c) Beskriv og implementer feilhåndtering på all I/O i oppgave a) og b). Se https://blog.golang.org/error-handling-and-go (Lenker til en ekstern side.)Lenker til en ekstern side. for mer informasjon om feilhåndtering i Go.

Vi har laget håndtering (se oppgave 3A og 3B) til at tilfellet der brukeren skriver inn en bokstav (f.eks a, b, c) istedenfor et tall (f.eks 1, 2, 1084), så returnerer programmet en feil. Deretter ber den brukeren om å legge inn en akseptabel verdi.

#### d) Implementer håndtering av SIGINT i programmene i a) og b); programmene skal skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullfører naturlig.

Både a) og b) har implementert SIGINT.

#### e) Bygg filene som i oppgave 1 og 2 til kjørbare filer på ditt operativsystem og legg dem ved i /bin mappen.

Oppgave 1: Den kompilerte filen (fileinfo.exe) kan bli funnet i (https://raw.githubusercontent.com/StavenX/IS-105_oppgaver/master/Oblig2/bin/fileinfo.exe).

Oppgave 2: Den kompilerte filen (filecount.exe) kan bli funnet i(https://raw.githubusercontent.com/StavenX/IS-105_oppgaver/master/Oblig2/bin/filecount.exe).

----------------------------------------------------------------------------------------------------------
