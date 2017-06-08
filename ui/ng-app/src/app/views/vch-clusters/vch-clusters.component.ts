import { Component, OnInit } from '@angular/core';
import { Links } from '../../utils/links';
import { DocumentService } from '../../utils/document.service';
import * as I18n from 'i18next';

@Component({
  selector: 'app-vch-clusters',
  templateUrl: './vch-clusters.component.html',
  styleUrls: ['./vch-clusters.component.scss']
})
export class VchClustersComponent implements OnInit {

  constructor(private service: DocumentService) { }

  serviceEndpoint = Links.CLUSTERS;

  ngOnInit() {
  }

  cpuPercentageLevel(cluster) {
    if (!cluster){
      return 0;
    }
    return Math.floor(cluster.cpuUsage / cluster.totalCpu * 100);
  }

  memoryPercentageLevel(cluster) {
    if (!cluster){
      return 0;
    }
    return Math.floor(cluster.memoryUsage / cluster.totalMemory * 100);
  }

  clusterState(cluster) {
    return I18n.t('clusters.state.' + cluster.status);
  }
}
