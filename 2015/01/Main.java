import java.io.*;

public class Main{
	public static void main(String args[]){
		int andar = 0;
		char sobe = "(".charAt(0);
		char desce = ")".charAt(0);

		String arquivo = "input.txt";
		try{
			BufferedReader br = new BufferedReader(new FileReader(arquivo));
			String linhas = br.readLine();
			for(int i = 0; i < linhas.length(); i++){
				char letra = linhas.charAt(i);
				if(letra == sobe){
					andar = andar + 1;
				}else if(letra == desce){
					andar = andar - 1;
				}
			}
			System.out.println("Resultado: "+andar);
		}catch(FileNotFoundException ex){
			ex.printStackTrace();
		}catch(IOException e){
			e.printStackTrace();
		}
	}
}
