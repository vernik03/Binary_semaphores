package com.lab3;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Random;
import java.util.concurrent.Semaphore;

import static java.lang.Thread.sleep;

public class Barbershop {

    Queue<Thread> clients;
    Barber barber;
    private Semaphore sem;

    Barbershop() {
        sem = new Semaphore(2);
        clients = new LinkedList<Thread>();
        barber = new Barber();
        new Thread(barber).start();

        for (int i = 0; i < 10; i++) {
            clients.add(new Thread(new Client()));
            clients.element().start();
            Random rand = new Random();
            int int_random = rand.nextInt(2000)+500;
            try {
                sleep(int_random);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    class Client implements Runnable{
        Client(){}

        public void run(){
            if (clients.size() == 1) {
                try {
                    sem.acquire();
                    System.out.println(" садится в кресло");
                    wait();
                    System.out.println("постригся");
                    sem.release();

                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    class Barber implements Runnable{

        Barber(){}
        public void run(){
            while (true) {
                try {
                sem.acquire();
                sleep(200);
                System.out.println (" начинает стричь");

                sleep(1000);
                notify();
                System.out.println ("постриг");
                sleep(200);
                clients.remove();
                sem.release();

                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    /*class Pot{
        private int value;
        private int max;
        Pot(){
            this.max = 10;
            value = 0;
        }

        public synchronized void addHoney(){
            while (value == this.max){
                try {
                    wait();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            value++;
            System.out.println("Bee added 1 honey. Pot value: " + value);
            notify();
        }

        public synchronized void clearPot(){
            while (value < this.max){
                try {
                    wait();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            value = 0;
            System.out.println("Bear ate all honey");
            notify();
        }
    }*/
}
