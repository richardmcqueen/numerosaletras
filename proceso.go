package numerosaletras


// convierte un numero as letras


func Convertir(n int) string {
  var respuesta string
  if n>= 10 {
  	respuesta = primerosDiez[n%10]
  } else {
  	respuesta = unidades[n]
  }

  return respuesta

}


